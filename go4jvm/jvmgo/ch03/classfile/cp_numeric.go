package classfile

import "math"

/*
Constant_Integer_Info{
u1 tag
u4 bytes;
}
*/

type ConstantIntegerInfo struct {
	val int32
}

func (self *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	// 如果是负数的常量?????
	byte := reader.readUint32()
	self.val = int32(byte)
}

func (self *ConstantIntegerInfo) Value() int32 {
	return self.val
}

/*
Constant_Float_Info{
u1 tag
u4 bytes;
}
*/

type ConstantFloatInfo struct {
	val float32
}

func (self *ConstantFloatInfo) readInfo(reader *ClassReader) {
	byte := reader.readUint32()
	self.val = math.Float32frombits(byte)
}

func (self *ConstantFloatInfo) Value() float32 {
	return self.val
}

/*
Constant_Long_Info{
u1 tag
u4 high_bytes
u4 low_bytes
}
*/

type ConstantLongInfo struct {
	val int64
}

func (self *ConstantLongInfo) readInfo(reader *ClassReader) {
	byte := reader.readUint64()
	self.val = int64(byte)
}

func (self *ConstantLongInfo) Value() int64 {
	return self.val
}
