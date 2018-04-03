package classfile

/*
CONSTANT_NameAndType_info{
u1 tag
u2 name_index
u2 descriptor_index
}

CONSTANT_NameAndType_info用于标识调用的成员字段或者方法；
类型描述符：
byte B
short S
char C
int I
long J
float F
double D
boolean Z
引用类型  L+类型限定名+;
数组类型 [+数组元素类型描述符

字段描述符就是字段类型的描述符
方法描述符是  (参数类型描述符)+返回值类型描述符 void是单个字母V
*/

type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
	self.descriptorIndex = reader.readUint16()
}
