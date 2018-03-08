package rtda

import "math"

/*
GO语言的int类型因平台不同而不同，64位是int64,32位是int32；
uintptr可以存放一个内存地址的整型；可以和unsafe.Pointer进行类型互转；
uintptr指向的变量还是可能被内存回收，但是unsafe.Pointer是不会被回收;

局部变量表是通过索引访问
*/
// 局部变量表
type LocalVars []Slot

func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}

//-----------------------------针对不同类型数据的处理
func (self LocalVars) SetInt(index uint, val int) {
	self[index].num = val
}

func (self LocalVars) GetInt(index uint) int {
	return self[index].num
}

/*
float通过先转换成bit，再从bit转换为int
*/
func (self LocalVars) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	//noinspection GoBinaryAndUnaryExpressionTypesCompatibility
	self[index].num = int32(bits)
}

func (self LocalVars) GetFloat(index uint) float32 {
	bits := uint32(self[index].num)
	return math.Float32frombits(bits)
}

//noinspection GoBinaryAndUnaryExpressionTypesCompatibility
func (self LocalVars) SetLong(index uint, val int64) {
	// 64位赋值32位，会直接截取后32位
	// 低位在前，高位在后
	self[index].num = int32(val)
	self[index+1].num = int32(val >> 32)
}

//noinspection GoBinaryAndUnaryExpressionTypesCompatibility
func (self LocalVars) GetLong(index uint) int64 {
	low := uint32(self[index].num)
	high := uint32(self[index+1].num)
	//32位转64位，直接放在低位
	return int64(high)<<32 | low
}
