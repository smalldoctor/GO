package classfile

/*
CONSTANT_String_info{
u1 tag
u2 string_index ==> cp_index
}
JVM规范中的字符串是java.lang.String的字符串常量；但是JVM规范中Class文件中的字符串不直接保存
字符串，而是指向常量池中一个CONSTANT_Utf8_info类型的常量；

"test"会生成一个CONSTANT_Utf8_info;如果"test"被使用，则会生成一个CONSTANT_String_info；
"test"不论被使用多少次，都公用一个CONSTANT_String_info；如果不被使用，则不会生成CONSTANT_String_info；
"test"不论是否被使用都会生成CONSTANT_Utf8_info；
*/

type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

func (self *ConstantStringInfo) readInfo(reader *ClassReader) {
	self.stringIndex = reader.readUint16()
}

/*
返回String常量
*/
func (self *ConstantStringInfo) String() string {
	return self.cp.getUtf8(self.stringIndex)
}
