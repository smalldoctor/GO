package rtda

/*
JVM依赖于系统实现线程，线程执行字节码；底层线程负责执行，JVM相当于中间层，上层
则是字节码;

通过-Xss设置栈的大小
*/

type Thread struct {
	// 程序计数器
	pc int
	//JVM栈
	stack *Stack
}

func NewThread() *Thread {
	return &Thread{
		// 默认frame大小为1024
		stack: newStack(1024),
	}
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
