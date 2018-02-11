package classfile

/*
因为常量池中表项的具体类型不确定，因此定义一个接口，根据实际情况读取字节流
*/
type ConstantInfo interface {
	readInfo(reader *ClassReader)
}
