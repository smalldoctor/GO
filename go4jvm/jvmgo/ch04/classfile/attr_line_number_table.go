package classfile

/*
LineNumberTable用来存放方法的行号
LocalVariableTable用来存放方法的局部变量信息
LineNumberTable，LocalVariableTable，SourceFile都属于调试用信息，不是运行时必须；
默认javac会生成这些信息，但是可以通过-g:none关闭这些信息；

LineNumberTable_attribute{
u2 attribute_name_index
u4 attribute_length
u2 line_number_table_length
{
u2 start_pc
u2 line_number
}line_number_table[line_number_table_length]
}
*/

type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}
type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (self *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	self.lineNumberTable = readLineNumberTableEntry(reader)
}
func readLineNumberTableEntry(reader *ClassReader) []*LineNumberTableEntry {
	lineNumberTableLength := reader.readUint16()
	lineNumberTable := make([]*LineNumberTableEntry, lineNumberTableLength)
	for i := range lineNumberTable {
		lineNumberTable[i] = &LineNumberTableEntry{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
	return lineNumberTable
}
