package constants

import (
	"go4jvm/jvmgo/ch05/instructions/base"
	"go4jvm/jvmgo/ch05/rtda"
)

/*将byte和short类型的操作数转换为int型压入操作数栈*/

// byte
type BIPUSH struct {
	val int8
}

// short
type SIPUSH struct {
	val int16
}

// 获取byte型操作数（8位）
func (self *BIPUSH) FetchOperand(reader *base.ByteCodeReader) {
	self.val = reader.Readint8()
}
func (self *BIPUSH) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(int32(self.val))
}

// 获取short型操作数（16位）
func (self *SIPUSH) FetchOperand(read *base.ByteCodeReader) {
	self.val = read.ReadInt16()
}
func (self *SIPUSH) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(int32(self.val))
}
