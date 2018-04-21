package base

/*
读取字节码获取操作数
*/
type ByteCodeReader struct {
	// 存放字节码
	code []byte
	// pc字段记录读取到哪个字节；指向当前可以被读取的字节
	pc int
}

func (self *ByteCodeReader) ReadUint8() uint8 {
	code := self.code[self.pc]
	self.pc++
	return code
}

func (self *ByteCodeReader) Readint8() int8 {
	return int8(self.ReadUint8())
}

func (self *ByteCodeReader) ReadUint16() uint16 {
	code1 := uint16(self.ReadUint8())
	code2 := uint16(self.ReadUint8())
	return code1<<8 | code2
}

func (self *ByteCodeReader) ReadInt16() int16 {
	return int16(self.ReadUint16())
}

/*
在存储字节码时，是大端；
*/
func (self *ByteCodeReader) ReadInt32() int32 {
	code1 := int32(self.ReadUint8())
	code2 := int32(self.ReadUint8())
	code3 := int32(self.ReadUint8())
	code4 := int32(self.ReadUint8())
	return code1<<24 | code2<<16 | code3<<8 | code4
}
