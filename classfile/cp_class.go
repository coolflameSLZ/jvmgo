package classfile

/**
class name的信息，也出现在常量池中的utf8中。
同样是根据索引在mutf8中找
*/
type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}

func (self *ConstantClassInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.ReadUint16()
}
func (self *ConstantClassInfo) Name() string {
	return self.cp.GetUtf8(self.nameIndex)
}
