package stack

import (
	"go4jvm/jvmgo/ch05/instructions/base"
	"go4jvm/jvmgo/ch05/rtda"
)

/*
栈顶元素复制
*/
type DUP struct {
	base.NoOperandsInstruction
}

/*
复制栈顶元素，压入栈
*/
func (self *DUP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot1)
}

type DUP_X1 struct {
	base.NoOperandsInstruction
}

/*
bottom -> top
[...][c][b][a]
          __/
         |
         V
[...][c][a][b][a]
*/
func (self *DUP_X1) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

type DUP_X2 struct {
	base.NoOperandsInstruction
}

/*
bottom -> top
[...][c][b][a]
       _____/
      |
      V
[...][a][c][b][a]
*/
func (self *DUP_X2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

type DUP2 struct {
	base.NoOperandsInstruction
}

/*
bottom -> top
[...][c][b][a]____
          \____   |
               |  |
               V  V
[...][c][b][a][b][a]
*/
func (self *DUP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slota := stack.PopSlot()
	slotb := stack.PopSlot()
	stack.PushSlot(slotb)
	stack.PushSlot(slota)
	stack.PushSlot(slotb)
	stack.PushSlot(slota)
}

type DUP2_X1 struct {
	base.NoOperandsInstruction
}

/*
bottom -> top
[...][c][b][a]
       _/ __/
      |  |
      V  V
[...][b][a][c][b][a]
*/
func (self *DUP2_X1) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slota := stack.PopSlot()
	slotb := stack.PopSlot()
	slotc := stack.PopSlot()
	stack.PushSlot(slotb)
	stack.PushSlot(slota)
	stack.PushSlot(slotc)
	stack.PushSlot(slotb)
	stack.PushSlot(slota)
}

type DUP2_X2 struct {
	base.NoOperandsInstruction
}

/*
bottom -> top
[...][d][c][b][a]
       ____/ __/
      |   __/
      V  V
[...][b][a][d][c][b][a]
*/

func (self *DUP2_X2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slota := stack.PopSlot()
	slotb := stack.PopSlot()
	slotc := stack.PopSlot()
	slotd := stack.PopSlot()
	stack.PushSlot(slotb)
	stack.PushSlot(slota)
	stack.PushSlot(slotd)
	stack.PushSlot(slotc)
	stack.PushSlot(slotb)
	stack.PushSlot(slota)
}
