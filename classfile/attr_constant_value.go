package classfile

type ConstantValueAttribute struct {
	constantValueIndex uint16
}

func (c *ConstantValueAttribute) readInfo(reader *ClassReader) {
	c.constantValueIndex = reader.ReadUint16()
}

func (c *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return c.constantValueIndex
}
