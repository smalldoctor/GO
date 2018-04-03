package classfile

/*
CONSTANT_Class_info{
u1 tag
u2 name_index
}

类名，父类明，接口数组中的接口都是CONSTANT_Class_info;
然后CONSTANT_Class_info中的name_index指向CONSTANT_Utf8_info；
*/

type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}

func (self *ConstantClassInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
}

func (self *ConstantClassInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}
