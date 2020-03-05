package classfile

type UnparsedAttribute struct {
	name   string
	length uint32
	info   []byte
}

// 根据自己的长度，将信息解析至自己的info属性中
func (u *UnparsedAttribute) readInfo(reader *ClassReader) {
	u.info = reader.ReadBytes(u.length)
}
