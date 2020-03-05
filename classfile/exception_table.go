package classfile

type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

/**
将错误信息解析至 exception实体类
*/
func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := reader.ReadUint16()
	exceptionTables := make([]*ExceptionTableEntry, exceptionTableLength)

	for i := range exceptionTables {

		exceptionTables[i] = &ExceptionTableEntry{
			startPc:   reader.ReadUint16(),
			endPc:     reader.ReadUint16(),
			handlerPc: reader.ReadUint16(),
			catchType: reader.ReadUint16(),
		}
	}

	return exceptionTables
}
