package rtda

type Thread struct {
	pc    int    // 寄存器
	stack *Stack // 私有栈
}

func NewThread(stack *Stack) *Thread {
	return &Thread{stack: NewStack(1024)}
}
func (self *Thread) PC() int {
	return self.pc
}
func (self *Thread) SetPC(pc int) {
	self.pc = pc
}

func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}
func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}

func (self *Thread) CurrentFrame() *Frame {
	return self.stack.top()
}
