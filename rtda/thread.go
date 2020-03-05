package rtda

type Thread struct {
	pc    int    // 寄存器
	stack *Stack // 私有栈
}

func NewThread(stack *Stack) *Thread {
	return &Thread{stack: NewStack(1024)}
}

func (t *Thread) Pc() int {
	return t.pc
}

func (t *Thread) SetPc(pc int) {
	t.pc = pc
}
func (t *Thread) PushFrame(frame *Frame) {
	t.stack.push(frame)
}
func (t *Thread) PopFrame() *Frame {
	t.stack.pop()
}
func (t *Thread) CurrentFrame() *Frame {
	t.stack.top()
}
