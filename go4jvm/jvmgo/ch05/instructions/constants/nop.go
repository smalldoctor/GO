package constants

import (
	"go4jvm/jvmgo/ch05/instructions/base"
	"go4jvm/jvmgo/ch05/rtda"
)

/*
常量指令21条；
空指令，什么都不做;
*/

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {
	//	空指令，什么都不做
}
