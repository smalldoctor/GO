package rtda

/*
JVM虚拟机栈中帧，保存方法的执行状态，包括局部变量表，操作数栈；
每个方法执行时对应一帧；
在编译阶段，class文件的code属性中已经确定了局部变量和操作数栈的大小；
*/

type Frame struct {
	// 通过链表实现栈;后继帧
	lower *Frame
	// 局部变量表指针
	localVars    LocalVars
	operandStack *OperandStack
	//	TODO 线程Thread
	nextPC int
}

/*
每个方法对应的局部变量表的大小和操作数栈的深度在编译阶段就已经由编译器计算好，
并且保存在字节码中，存在method_info的code属性中;
*/
func NewFrame(maxLocals, maxStacks uint) *Frame {
	return &Frame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStacks),
	}
}

// getters & setters
func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}
func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}
func (self *Frame) NextPC() int {
	return self.nextPC
}
func (self *Frame) SetNextPC(nextPC int) {
	self.nextPC = nextPC
}
