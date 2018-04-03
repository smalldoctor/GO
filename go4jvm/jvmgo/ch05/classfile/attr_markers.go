package classfile

/*
Deprecated_attribute{
u2 attribute_name_index
u4 attribute_length
}

Synthetic_attribute{
u2 attribute_name_index
u4 attribute_length
}

Deprecated和Synthetic属于标记型属性；属性内容为空，因此属性长度是0；
*/

type MarkerAttribute struct {
	// 根据JVM规范的定义，没有具体的内容
}

type DeprecatedAttribute struct {
	MarkerAttribute
}

type SyntheticAttribute struct {
	MarkerAttribute
}

func (self *MarkerAttribute) readInfo(reader *ClassReader) {
	// 根据JVM规范的定义，没有具体的内容
}

