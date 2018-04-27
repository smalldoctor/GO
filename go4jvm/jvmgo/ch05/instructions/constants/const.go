package constants

import (
	"go4jvm/jvmgo/ch05/instructions/base"
	"go4jvm/jvmgo/ch05/rtda"
)

/*
const常量指令
*/

// NULL引用推入操作数栈
type ACONST_NULL struct {
	base.NoOperandsInstruction
}

//把double型0推入操作数栈
type DCONST_0 struct {
	base.NoOperandsInstruction
}

type DCONST_1 struct {
	base.NoOperandsInstruction
}

type FCONST_0 struct {
	base.NoOperandsInstruction
}

type FCONST_1 struct {
	base.NoOperandsInstruction
}

type FCONST_2 struct {
	base.NoOperandsInstruction
}

// 把int型-1压入操作数栈
type ICONST_M1 struct {
	base.NoOperandsInstruction
}

type ICONST_0 struct {
	base.NoOperandsInstruction
}

type ICONST_1 struct {
	base.NoOperandsInstruction
}

type ICONST_2 struct {
	base.NoOperandsInstruction
}

type ICONST_3 struct {
	base.NoOperandsInstruction
}

type ICONST_4 struct {
	base.NoOperandsInstruction
}

type ICONST_5 struct {
	base.NoOperandsInstruction
}

type LCONST_0 struct {
	base.NoOperandsInstruction
}

type LCONST_1 struct {
	base.NoOperandsInstruction
}

func (self *ACONST_NULL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushRef(nil)
}

func (self *DCONST_0) Execute(frame *rtda.Frame) {
	// double类型 0
	frame.OperandStack().PushDouble(0.0)
}

func (self *ICONST_M1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(-1)
}
