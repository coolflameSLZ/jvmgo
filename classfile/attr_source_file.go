package classfile

// 只需要记录 sourceFile所在的索引就好，
type SourceFileAttribute struct {
	cp              ConstantPool
	sourceFileIndex uint16
}

func (s *SourceFileAttribute) readInfo(reader *ClassReader) {
	s.sourceFileIndex = reader.ReadUint16()
}

func (s *SourceFileAttribute) FileName() string {
	return s.cp.GetUtf8(s.sourceFileIndex)
}
