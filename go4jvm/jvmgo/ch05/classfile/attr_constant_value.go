package classfile

/*
ConstantValue_attribute{
u2 attribute_name_index
u4 attribute_length
u2 constantvalue_index
}

ConstantValue_attribute常量属性，定长属性，只会出现在Field_info中，表示常量表达式的值;
constantvalue_index指向常量池的索引；
attribute_length长度必须是2；

int,byte,short,char,boolean ==》 CONSTANT_Integer_info
String ==> CONSTANT_String_info
*/

type ConstantValueAttribute struct {
	constantValueIndex uint16
}

func (self *ConstantValueAttribute) readInfo(reader *ClassReader) {
	self.constantValueIndex = reader.readUint16()
}

func (self *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return self.constantValueIndex
}
