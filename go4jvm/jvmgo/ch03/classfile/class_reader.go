package classfile

import "encoding/binary"

/*
定义一个结构体包含需要读取的class文件的的字节流，定义读取各种类型的方法；

文件结构：
ClassFile{
u4 magic;
u2 minor_version;
u2 major_version;
u2 constant_pool_count;
cp_info constant_pool[constant_pool_count - 1];  //??为什么-1
u2 access_flags;
u2 this_class;
u2 super_class;
u2 interfaces_count;
u2 interfaces[interfaces_count];
u2 fields_count;
field_info fields[fields_count];
u2 methods_count;
method_info methods[methods_count];
u2 attributes_count;
attribute_info attributes[attributes_count];
}
*/
type ClassReader struct {
	data []byte
}

func (self *ClassReader) readUint8() uint8 {
	// 无符号8位
	val := self.data[0]
	self.data = self.data[1:]
	return val
}

/*
binary.BigEndian 库用来操作字节，如大位，小位字节等
*/
func (self *ClassReader) readUint16() uint16 {
	/*
	不通过索引记录读取的位置，通过reslice的方式跳过已读取的数据
	*/
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return val
}

func (self *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return val
}

/*
JVM规范中没有64位的
*/
func (self *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return val
}

/*
表头是u2类型的；每个表项也是u2类型;
*/
func (self *ClassReader) readUint16s() []uint16 {
	n := self.readUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = self.readUint16()
	}
	return s
}

//读取指定数量的字节数
func (self *ClassReader) readBytes(n uint32) []byte {
	val := self.data[:n]
	self.data = self.data[n:]
	return val
}
