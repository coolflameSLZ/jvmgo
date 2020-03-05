package classfile

/**
Code属性中存放字节码等方法相关信息
*/
type CodeAttribute struct {
	cp             ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributes     []AttributeInfo
}

func (c *CodeAttribute) readInfo(reader *ClassReader) {

	c.maxStack = reader.ReadUint16()
	c.maxLocals = reader.ReadUint16()
	codeLength := reader.ReadUint32()
	c.code = reader.ReadBytes(codeLength)
	c.exceptionTable = readExceptionTable(reader)
	c.attributes = ReadAttributes(reader, c.cp)
}
