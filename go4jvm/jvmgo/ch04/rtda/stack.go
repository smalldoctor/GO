package rtda

/*
如果stack存在大小限制,栈的深度超出，则导致StackOverflowError；
如果stack动态扩展时，没有足够的空间，则抛出OutOfMemoryError;
通过链表的方式实现栈，可以做到空间的动态分配；并且出栈之后，可以利用GO进行
垃圾回收
*/
type Stack struct {
	//当前栈帧数
	size uint
	/*
	每个JVM栈允许的帧的个数
	*/
	maxSize uint
	// 当前帧，栈顶
	_top *Frame
}

func newStack(maxSzie uint) *Stack {
	return &Stack{
		maxSize: maxSzie,
	}
}

func (self *Stack) push(frame *Frame) {
	// 如果帧已满则抛出StackOverflowError
	if self.size >= self.maxSize {
		panic("java.lang.StackOverflowError")
	}
	if self._top != nil {
		frame.lower = self._top
	}
	self._top = frame
	self.size++
}

func (self *Stack) pop() *Frame {
	//  出栈
	if self._top == nil {
		panic("jvm stack is empty")
	}
	top := self._top
	self._top = top.lower
	top.lower = nil
	self.size--
	return top
}

func (self *Stack) top() *Frame {
	//  当前帧,但是不出栈
	if self._top == nil {
		panic("jvm stack is empty")
	}
	return self._top
}
