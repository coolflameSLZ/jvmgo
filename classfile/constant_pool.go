package classfile

type ConstantPool []ConstantInfo

// 从 class 文件中提取出ConstantPool常量池信息
func ReadConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.ReadUint16())
	cp := make([]ConstantInfo, cpCount)

	for i := 1; i < cpCount; i++ {
		cp[i] = readConstantInfo(reader, cp)

		// 有的类型需要占两个位置
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}

	return cp
}

// 按照索引查找常量
func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := self[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index ,不合法的常量池数组")
}

// 从常量池中查找字段或者方法的名字和描述符
func (self ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := self.GetUtf8(ntInfo.nameIndex)
	_type := self.GetUtf8(ntInfo.descriptorIndex)
	return name, _type
}

// 从常量池查找类名
func (self ConstantPool) getClassName(index uint16) string {
	classInfo := self.getConstantInfo(index).(*ConstantClassInfo)
	return self.GetUtf8(classInfo.nameIndex)
}

//从常量池查找UTF-8字符串
func (self ConstantPool) GetUtf8(index uint16) string {
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}
