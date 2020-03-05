package classfile

const (
	CONSTANT_Class              = 7
	CONSTANT_Fieldref           = 9
	CONSTANT_Methodref          = 10
	CONSTANT_InterfaceMethodref = 11
	CONSTANT_String             = 8
	CONSTANT_Integer            = 3
	CONSTANT_Float              = 4
	CONSTANT_Long               = 5
	CONSTANT_Double             = 6
	CONSTANT_NameAndType        = 12
	CONSTANT_Utf8               = 1
	CONSTANT_MethodHandle       = 15
	CONSTANT_MethodType         = 16
	CONSTANT_InvokeDynamic      = 18
)

/**
readInfo（）方法读取常量信息，需要由具体的常量结构体实现。

readConstantInfo（）函数先读出tag值，
然后调用newConstantInfo（）函数创建具体的常量，
最后调用常量的readInfo（）方法读取常量信息，
*/

type ConstantInfo interface {
	readInfo(reader *ClassReader)
}

// 从常量池中解析成常量子类型
func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	tag := reader.ReadUint8()

	constantInfo := newConstantInfo(tag, cp)
	constantInfo.readInfo(reader)
	return constantInfo
}

// 构造函数, tag 是常量类型标识符,根据不同的tag值，构造不同的ConstantInfo
func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch tag {
	case CONSTANT_Integer:
		return &ConstantIntegerInfo{}
	case CONSTANT_Float:
		return &ConstantFloatInfo{}
	case CONSTANT_Long:
		return &ConstantLongInfo{}
	case CONSTANT_Double:
		return &ConstantDoubleInfo{}
	case CONSTANT_Utf8:
		return &ConstantUtf8Info{}
	case CONSTANT_String:
		return &ConstantStringInfo{cp: cp}
	case CONSTANT_Class:
		return &ConstantClassInfo{cp: cp}
	case CONSTANT_Fieldref:
		return &ConstantFieldrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_Methodref:
		return &ConstantMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_InterfaceMethodref:
		return &ConstantInterfaceMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_NameAndType:
		return &ConstantNameAndTypeInfo{}
	//	下面的是jvm se 7开始实现的东西，暂不考虑。
	//case CONSTANT_MethodType:
	//	return &ConstantMethodTypeInfo{}
	//case CONSTANT_MethodHandle:
	//	return &ConstantMethodHandleInfo{}
	//case CONSTANT_InvokeDynamic:
	//	return &ConstantInvokeDynamicInfo{}
	default:
		panic("java.lang.ClassFormatError: constant pool tag! ,这个tag值不是标注的常量值类型,tat=" + string(tag))
	}
}
