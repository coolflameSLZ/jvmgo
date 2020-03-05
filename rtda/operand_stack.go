package rtda

import (
	"math"
)

type OperandStack struct {
	size  uint //用于记录栈顶位置
	slots []Slot
}

func newOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{
			slots: make([]Slot, maxStack),
		}
	}
	return nil
}

func (self *OperandStack) PushInt(val int32) {
	self.slots[self.size].num = val
	self.size++
}
func (self *OperandStack) PopInt() int32 {
	self.size--
	return self.slots[self.size].num
}

/**
float变量还是先
转成int类型，然后按int变量处理
*/
func (self *OperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	self.slots[self.size].num = int32(bits)
	self.size++
}
func (self *OperandStack) PopFloat() float32 {
	self.size--
	bits := uint32(self.slots[self.size].num)
	return math.Float32frombits(bits)
}

/**
把long变量推入栈顶时，要拆成两个int变量。弹出时，先弹出
两个int变量，然后组装成一个long变量
*/
func (self *OperandStack) PushLong(val int64) {
	self.slots[self.size].num = int32(val)
	self.slots[self.size+1].num = int32(val >> 32)
	self.size += 2
}
func (self *OperandStack) PopLong() int64 {
	self.size -= 2
	low := uint32(self.slots[self.size].num)
	high := uint32(self.slots[self.size+1].num)
	return int64(high)<<32 | int64(low)
}

/**
double变量先转成long类型，然后按long变量处理
*/
func (self *OperandStack) PushDouble(val float64) {
	bits := math.Float64bits(val)
	self.PushLong(int64(bits))
}
func (self *OperandStack) PopDouble() float64 {
	bits := uint64(self.PopLong())
	return math.Float64frombits(bits)
}

/**
引用类型操作起来同int
弹出引用后，把Slot结构体的ref字段设置成nil，
这样做是为了帮助Go的垃圾收集器回收Object结构体实例
*/
func (self *OperandStack) PushRef(ref *Object) {
	self.slots[self.size].ref = ref
	self.size++
}
func (self *OperandStack) PopRef() *Object {
	self.size--
	ref := self.slots[self.size].ref
	self.slots[self.size].ref = nil
	return ref
}
