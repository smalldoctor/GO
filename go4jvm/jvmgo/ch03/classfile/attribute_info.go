package classfile

/*
属性可以存放各种信息，每个虚拟机可以自定义进行扩展；java虚拟机会跳过自己不认识的属性
通过属性名标识不同的属性；
attribute_info{
u2 attribute_name_index 指向常量池的CONSTANT_Utf8_info
u4 attribute_length u4长度，即32位的二进制整数
u1 info[attribute_length]
// 接下来的N个字节组成具体的内容；不同的属性定义N个字节的每个字节的含义不同；
}

JVM总共定义了三组23个属性；
第一组5个JVM必须的；
第二组12个java类库所必须的；
第三组6个可选的，如果class文件中出现这些属性，JVM和Java类库可以使用；
*/

type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	attributeCount := reader.readUint16()
	attributes := make([]AttributeInfo, attributeCount)
	for i := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}
	return attributes
}

func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	attrNameIndex := reader.readUint16()
	// 因为属性是按照名字进行区分的，所以需要根据索引读取具体的名字
	attrName := cp.getUtf8(attrNameIndex)
	attrLen := reader.readUint32()
	attrInfo := newAttributeInfo(attrName, attrLen, cp)
	attrInfo.readInfo(reader)
	return attrInfo
}

func newAttributeInfo(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {
	// 按照不同的属性名不同的定义构建属性信息
	switch attrName {
	case "Code":
		return &CodeAttribute{cp: cp}
	case "ConstantValue":
		return &ConstantValueAttribute{}
	case "Deprecated":
		return &DeprecatedAttribute{}
	case "Exceptions":
		return &ExceptionsAttribute{}
	case "LineNumberTable":
		return &LineNumberTableAttribute{}
	case "LocalVariableTable":
		return &LocalVariableTableAttribute{}
	case "SourceFile":
		return &SourceFileAttribute{cp: cp}
	case "Synthetic":
		return &SyntheticAttribute{}
	default:
		// 默认忽视不能识别的属性表
		return &UnparsedAttribute{attrName, attrLen, nil}
	}
}
