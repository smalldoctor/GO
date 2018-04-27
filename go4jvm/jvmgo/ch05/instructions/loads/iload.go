package loads

import (
	"go4jvm/jvmgo/ch05/instructions/base"
	"go4jvm/jvmgo/ch05/rtda"
)

/*加载指令，从局部变量获取数据，并压入操作数栈;共33条指令
根据指令操作变量的不同，分为6类：
引用类型    aload
double     dload
float      fload
int        iload
long       lload
数组        xaload
*/

/*
获取操作数（即局部变量表的索引）对应的变量，推入操作数栈
*/
type ILOAD struct {
	base.Index8Instruction
}

func (self *ILOAD) Execute(frame *rtda.Frame) {
	_iLoad(frame, self.Index)
}

type ILOAD_0 struct {
	base.NoOperandsInstruction
}

func (self *ILOAD_0) Execute(frame *rtda.Frame) {
	_iLoad(frame, 0)
}

type ILOAD_1 struct {
	base.NoOperandsInstruction
}

func (self *ILOAD_1) Execute(frame *rtda.Frame) {
	_iLoad(frame, 0)
}

type ILOAD_2 struct {
	base.NoOperandsInstruction
}

func (self *ILOAD_2) Execute(frame *rtda.Frame) {
	_iLoad(frame, 0)
}

type ILOAD_3 struct {
	base.NoOperandsInstruction
}

func (self *ILOAD_3) Execute(frame *rtda.Frame) {
	_iLoad(frame, 0)
}

/*
避免重复代码，可以抽取为函数
*/
func _iLoad(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}
