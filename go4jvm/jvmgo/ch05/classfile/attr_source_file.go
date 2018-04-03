package classfile

/*
SourceFile_attribute{
u2 attribute_name_index
u4 attribute_length
u2 sourcefile_index
}

SourceFile_attribute出现在ClassFile中，用于指定源文件名，定长属性；
sourcefile_index指向CONSTANT_Utf8_info常量;
attribute_length必须是2，因为常量池是u2类型，即2个字节；
*/

type SourceFileAttribute struct {
	cp              ConstantPool
	sourceFileIndex uint16
}

func (self *SourceFileAttribute) readInfo(reader *ClassReader) {
	self.sourceFileIndex = reader.readUint16()
}

func (self *SourceFileAttribute) FileName() string {
	return self.cp.getUtf8(self.sourceFileIndex)
}
