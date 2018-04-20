package base

import "go4jvm/jvmgo/ch04/rtda"

/*
定义一个接口类，然后各种结构体实现接口;
interface <= abstract <=  specific
指令的执行过程分为：
1. 计算pc
2. 指令解码，包括获取操作数
3. 执行指令

总共205条指令，11个大类；保留指令3条；
*/
type Instruction interface {
	/*
	获取操作数
	*/
	FetchOperands(reader *ByteCodeReader)
	/*
	执行指令
	*/
	Execute(frame *rtda.Frame)
}

// 无操作数指令
type NoOperandsInstruction struct {
}

func (self *NoOperandsInstruction) FetchOperands(reader *ByteCodeReader) {
	// nothing to do
}

// 跳转指令
type BranchInstruction struct {
	offset int
}

func (self *BranchInstruction) FetchOperands(reader *ByteCodeReader) {
	self.offset = int(reader.ReadInt16())
}

type Index8Instruction struct {
	Index uint
}

func (self *Index8Instruction) FetchOperands(reader *ByteCodeReader) {
	self.Index = uint(reader.ReadUint8())
}

type Index16Instruction struct {
	Index uint
}

func (self *Index16Instruction) FetchOperands(reader *ByteCodeReader) {
	self.Index = uint(reader.ReadUint16())
}
