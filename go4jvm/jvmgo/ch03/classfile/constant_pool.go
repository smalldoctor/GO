package classfile

/*
常量池包含数字，字符串，字段名，方法名等各种各样的常量信息;
常量池在class文件中用表进行标识；
1. 表头是n，则表项个数为n-1，因为第0个是无效索引；
2. 表项的有效索引为1 ~ n-1, 索引0是无效索引,标示不指向任何常量
3. CONSTANT_Long_Info 和 CONSTANT_Double_Info 占据两个位置；因此实际表项的个数小于等与n-1个
*/
type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	// 常量池大小
	cpCount := reader.readUint16()
	// ConstantInfo 是interface，此处的实现者就是指针
	cp := make([]ConstantInfo, cpCount)
	var i uint16 = 1
	// 索引从1开始
	for i = 1; i < cpCount; i++ {
		// TODO 读取数据
		// TODO 如果是long,double需要占据两个位置
	}

	return cp
}

/*
获取指定索引的常量
*/
func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := self[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index!")
}

/*
从常量池中查找方法或者字段的名称和描述(类型)
*/
func (self ConstantPool) getNameAndType(index uint16) (string, string) {
	//TODO
	return "", ""
}

/*
从常量池中获取类信息
*/
func (self ConstantPool) getClassName(index uint16) string {
	// TODO
	return ""
}

/*
获取UTF8字符
*/
func (self ConstantPool) getUtf8(index uint16) string {
	// TODO
	return ""
}
