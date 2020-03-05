package rtda

type Stack struct {
	maxSize uint
	size    uint
	_top    *Frame // 栈顶的Frame
}

func NewStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

/**
插入节点
1. self.next  =  head.next
2. head .next  =  self
*/
func (s *Stack) push(frame *Frame) {

	if s.size > s.maxSize {
		panic("java.lang.StackOverflowError,栈溢出")
	}

	// 移动 next指针
	if s._top != nil {
		frame.lower = s._top
	}
	s._top = frame
	s.size++
}

/**
tmp = head
head.next = head.next.next
head.next =  nil
*/
func (self *Stack) pop() *Frame {
	if self._top == nil {
		panic("jvm stack is empty")
	}
	self.size--
	// head.next
	top := self._top
	// head.next = head.next.next
	self._top = top.lower
	// 切断弹出元素的指针
	top.lower = nil
	return top
}

/**
只返回，不弹出
*/
func (self *Stack) top() *Frame {
	if self._top == nil {
		panic("jvm stack is empty")
	}
	return self._top
}
