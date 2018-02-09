package classfile

type ClassFile struct {
	magic        uint32
	minorVersion uint16
	majorVersion uint16
	// 常量池
	accessFlag uint16
	thisClass  uint16
	superClass uint16
	interfaces []uint16
	// fields
	// methods
	// attributes
}
