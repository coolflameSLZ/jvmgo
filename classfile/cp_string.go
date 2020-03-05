package classfile

/**
CONSTANT_String_info本身并不存放字符串数据，
只存了常量池索引，这个索引指向一个CONSTANT_Utf8_info常量
*/

type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

func (c *ConstantStringInfo) readInfo(reader *ClassReader) {
	c.stringIndex = reader.ReadUint16()
}

// 根据索引找到字符串
func (c *ConstantStringInfo) string() string {
	return c.cp.GetUtf8(c.stringIndex)
}
