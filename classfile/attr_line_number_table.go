package classfile

type LineNumberTableAttribute struct {
	LineNumberTable []*LineNumberEntry
}

type LineNumberEntry struct {
	strartPc   uint16
	lineNumber uint16
}

func (l LineNumberTableAttribute) readInfo(reader *ClassReader) {

	lineNumberTableLength := reader.ReadUint16()
	l.LineNumberTable = make([]*LineNumberEntry, lineNumberTableLength)

	for i := range l.LineNumberTable {
		l.LineNumberTable[i] = &LineNumberEntry{
			strartPc:   reader.ReadUint16(),
			lineNumber: reader.ReadUint16(),
		}
	}
}
