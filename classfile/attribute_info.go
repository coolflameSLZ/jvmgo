package classfile

/**
定义属性
*/

type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

// 根据名称和长度，在常量池中抽取attributeInfo
func newAttributeInfo(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {
	switch attrName {
	case "Code":
		return &CodeAttribute{cp: cp}
	case "ConstantValue":
		return &ConstantValueAttribute{}
	case "Deprecated":
		return &DeprecatedAttribute{}
	case "Exceptions":
		return &ExceptionsAttribute{}
	case "LineNumberTable":
		return &LineNumberTableAttribute{}
	case "LocalVariableTable":
		return &LocalVariableTableAttribute{}
	case "SourceFile":
		return &SourceFileAttribute{cp: cp}
	case "Synthetic":
		return &SyntheticAttribute{}
	default:
		return &UnparsedAttribute{attrName, attrLen, nil}
	}
}

// 读取多个属性
func ReadAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	attributesCount := reader.ReadUint16()
	attributes := make([]AttributeInfo, attributesCount)

	for i := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}
	return attributes
}

// 解析单个属性
func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	attrNameIndex := reader.ReadUint16()
	attrName := cp.GetUtf8(attrNameIndex)
	attrLen := reader.ReadUint32()
	attrInfo := newAttributeInfo(attrName, attrLen, cp)
	attrInfo.readInfo(reader)
	return attrInfo
}
