package classfile

import (
	"math"
)

/**
从常量池中读取数字，
读取的有：
integer
float
long
Double
*/

type ConstantIntegerInfo struct {
	val uint32
}
type ConstantFloatInfo struct {
	val float32
}
type ConstantLongInfo struct {
	val int64
}
type ConstantDoubleInfo struct {
	val float64
}

func (c *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	readUint32 := reader.ReadUint32()
	c.val = readUint32
}

func (self *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.ReadUint32()
	self.val = math.Float32frombits(bytes)
}

func (self *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.ReadUint64()
	self.val = int64(bytes)
}
func (self *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.ReadUint64()
	self.val = math.Float64frombits(bytes)
}
