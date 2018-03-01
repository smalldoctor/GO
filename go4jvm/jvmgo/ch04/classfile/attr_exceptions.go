package classfile

/*
Exceptions_attribute{
u2 attribute_name_index
u4 attribute_length
u2 number_of_exceptions
u2 exception_index_table[number_of_exceptions]
}
Exceptions_attribute 用来记录方法throws的异常；变长属性；
exception_index_table是常量池的CONSTANT_Class_info的数组；
*/

type ExceptionsAttribute struct {
	exceptionIndexTable []uint16
}

func (self *ExceptionsAttribute) readInfo(reader *ClassReader) {
	self.exceptionIndexTable = reader.readUint16s()
}

func (self *ExceptionsAttribute) ExceptionIndexTable() []uint16 {
	return self.exceptionIndexTable
}
