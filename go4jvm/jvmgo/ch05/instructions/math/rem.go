package math

import (
	"go4jvm/jvmgo/ch05/instructions/base"
	"go4jvm/jvmgo/ch05/rtda"
	"math"
)

/*
求余指令4条;
求余指令是将栈顶的两个值进行求余；
bottom -> top
[...][c][b][a]
b % a
*/
type DREM struct {
	base.NoOperandsInstruction
}

type FREM struct {
	base.NoOperandsInstruction
}

type IREM struct {
	base.NoOperandsInstruction
}

type LREM struct {
	base.NoOperandsInstruction
}

func (self *DREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slota := stack.PopDouble()
	slotb := stack.PopDouble()
	result := math.Mod(slotb, slota)
	stack.PushDouble(result)
}

func (self *FREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slota := stack.PopFloat()
	slotb := stack.PopFloat()
	// 转换成64进行求余，然后再转换为32
	result := float32(math.Mod(float64(slotb), float64(slota)))
	stack.PushFloat(result)
}

func (self *IREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slota := stack.PopInt()
	slotb := stack.PopInt()
	if slota == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	result := slotb % slota
	stack.PushInt(result)
}

func (self *LREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slota := stack.PopLong()
	slotb := stack.PopLong()
	if slota == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	result := slotb % slota
	stack.PushLong(result)
}
