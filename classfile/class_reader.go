package classfile

import "encoding/binary"

/**
读取类,使用该类可以读取 class 文件。
*/
type ClassReader struct {
	data []byte
}

// 读取一个uint8
func (self *ClassReader) ReadUint8() uint8 {
	val := self.data[0]
	self.data = self.data[1:]
	return val
}

// 读取一个uint16
func (self *ClassReader) ReadUint16() uint16 {
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return val
}

// 读取一个uint32
func (self *ClassReader) ReadUint32() uint32 {
	val := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return val
}

// 读取一个uint64
func (self *ClassReader) ReadUint64() uint64 {
	val := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return val
}

// 读取uint16表，表的大小由开头的uint16数据指出
func (self *ClassReader) ReadUint16s() []uint16 {
	n := self.ReadUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = self.ReadUint16()
	}
	return s
}

// 读取指定数量的字节
func (self *ClassReader) ReadBytes(n uint32) []byte {
	bytes := self.data[:n]
	self.data = self.data[n:]
	return bytes
}
