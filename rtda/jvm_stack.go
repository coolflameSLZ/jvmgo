package rtda

type Stack struct {
	maxSize uint
	size    uint
	_top    *Frame // 栈顶的Frame
}

func NewStack(maxSize uint) *Stack {
	return &Stack{maxSize: maxSize}
}

func (s *Stack) push(frame *Frame) {

	if s.size > s.maxSize {
		panic("java.lang.StackOverflowError,栈溢出")
	}

	if s._top != nil {
		frame.lower = s._top
	}

	s._top = frame
	s.size++
}
