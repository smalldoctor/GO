package classfile

/*
field_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
method_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
虚拟机规范定义的字段和方法的基本结构相同，仅在属性表存在差异；

method:
class文件在生成时，编译器会默认生成init方法,即默认构造器
*/

//字段和方法使用同一个结构体
type MemberInfo struct {
	//常量池指针,用于后面获取信息
	cp               ConstantPool
	accessFlags      uint16
	nameIndex        uint16
	descriptionIndex uint16
	// 属性表
	attributes []AttributeInfo
}

func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:               cp,
		accessFlags:      reader.readUint16(),
		nameIndex:        reader.readUint16(),
		descriptionIndex: reader.readUint16(),
		attributes:       readAttributes(reader, cp),
	}
}

func (self *MemberInfo) AccessFlags() uint16 {
	return self.accessFlags
}

// Name名字
func (self *MemberInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}

// Description描述
func (self *MemberInfo) Description() string {
	return self.cp.getUtf8(self.descriptionIndex)
}
