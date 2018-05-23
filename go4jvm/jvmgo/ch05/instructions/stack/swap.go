package stack

import (
	"go4jvm/jvmgo/ch05/instructions/base"
	"go4jvm/jvmgo/ch05/rtda"
)

type SWAP struct {
	base.NoOperandsInstruction
}

func (self *SWAP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slota := stack.PopSlot()
	slotb := stack.PopSlot()
	stack.PushSlot(slota)
	stack.PushSlot(slotb)
}
