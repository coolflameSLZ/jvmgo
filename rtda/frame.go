package rtda

type Frame struct {
	lower        *Frame        // 指向下一个栈帧，相当于 链表中的next
	localVars    LocalVars     // 保存局部变量的指针
	operandStack *OperandStack // 操作数栈的指针
}

func NewFrame(maxLoacl, maxStack uint) *Frame {
	return &Frame{
		localVars:    newLocalVars(maxLoacl),
		operandStack: newOperandStack(maxStack),
	}
}

// getters
func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}
func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}
