package rtda

/*
JVM虚拟机栈中帧，保存方法的执行状态，包括局部变量表，操作数栈；
每个方法执行时对应一帧；
*/

type Frame struct {
	// 通过链表实现栈;后继帧
	lower *Frame
	// 局部变量表指针
	localVars    LocalVars
	operandStack *OperandStack
}

/*
每个方法对应的局部变量表的大小和操作数栈的深度在编译阶段就已经由编译器计算好，
并且保存在字节码中，存在method_info的code属性中;
*/
func NewFrame(maxLocals, maxStacks uint) *Frame {
	return &Frame{localVars: newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStacks),
	}
}
