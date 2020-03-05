package classfile

type ExceptionsAttribute struct {
	exceptionIndexTable []uint16
}

func (e *ExceptionsAttribute) readInfo(reader *ClassReader) {
	e.exceptionIndexTable = reader.ReadUint16s()

}

func (e *ExceptionsAttribute) ExceptionIndexTable() []uint16 {
	return e.exceptionIndexTable
}
