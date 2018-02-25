package classfile

import "fmt"

/*
定义一个结构体包含需要读取的class文件的的字节流，定义读取各种类型的方法；
解析class文件的关键是每个字节都被读取到；

文件结构：
ClassFile{
u4 magic;
u2 minor_version;
u2 major_version;
u2 constant_pool_count;
//实际大小还是n，指向索引0不指向任何数据，所以个数是n-1
cp_info constant_pool[constant_pool_count - 1];
u2 access_flags;
u2 this_class;
u2 super_class;
u2 interfaces_count;
u2 interfaces[interfaces_count];
u2 fields_count;
field_info fields[fields_count];
u2 methods_count;
method_info methods[methods_count];
u2 attributes_count;
attribute_info attributes[attributes_count];
}
*/
type ClassFile struct {
	//magic        uint32
	minorVersion uint16
	majorVersion uint16
	// constant pool
	constantPool ConstantPool
	/*
	accessFlag是一个16的bitmask，包含多个信息
	*/
	accessFlags uint16
	thisClass   uint16
	superClass  uint16
	// 在字节码文件中存放的常量池的索引（CONSTANT_Class_info）
	interfaces []uint16
	fields     []*MemberInfo
	methods    []*MemberInfo
	// attributes
	// AttributeInfo是接口，实现者都是结构体指针
	attributes []AttributeInfo
}

func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		/*
		recover如在defer中执行，用来阻止同一个goroutine的panic；如果在defer外面执行，
		则无法阻止；
		recover如果在defer外面执行，或者未发生panic，或者发生panic时传递给panic的值为nil
		，则recover的值为nil；
		*/
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

/*
ClassFile作为目标结构，使用一个输入流作为参数;
因为最终操作的是字节流，所以需要搞清楚每个字节代表的含义；
*/
func (self *ClassFile) read(reader *ClassReader) {
	// 获取字节流中相关字节，将字节流转换为 ClassFile结构体
	self.readAndCheckMagic(reader)
	self.readAndCheckVersion(reader)
	// 常量池
	self.constantPool = readConstantPool(reader)
	self.accessFlags = reader.readUint16()
	self.thisClass = reader.readUint16()
	self.superClass = reader.readUint16()
	self.interfaces = reader.readUint16s()
	self.fields = readMembers(reader, self.constantPool)
	self.methods = readMembers(reader, self.constantPool)
	self.attributes = readAttributes(reader, self.constantPool)
}

/*
很多文件的格式要求文件必须以固定几个字节开头，用于识别文件类型，被称作为魔数;
class文件要求开头四个字节是 0xCAFEBABE
*/
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		// jvm在文件格式不对时抛出 java.lang.ClassFormatError
		panic("java.lang.ClassFormatError:magic!")
	}
}

func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	// 从java1.2(45)之后，小版本号都是0
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	switch self.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0 {
			return
		}
	}

	// jvm遇到不支持的class版本时，抛出java.lang.UnsupportedClassVersionError
	panic("java.lang.UnsupportedClassVersionError!")
}

func (self *ClassFile) MinorVersion() uint16 {
	return self.minorVersion
}

func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}

func (self *ClassFile) AccessFlags() uint16 {
	return self.accessFlags
}

// Getter Field
func (self *ClassFile) Fields() []*MemberInfo {
	return self.fields
}

// Getter Method
func (self *ClassFile) Methods() []*MemberInfo {
	return self.methods
}

/*
本身类名，父类名，接口名都是在Constant Pool中存放的；
本身类名和父类名都必须是有效的常量池索引；class文件存储的类名都是将点换成斜线的全限定类，被称为
二进制名；除了java.lang.Object的父类是0之外，其他的java类都必须是有效的父类;
*/
// Getter ClassName
func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName(self.thisClass)
}

// Getter SuperClassName
func (self *ClassFile) SuperClassName() string {
	return self.constantPool.getClassName(self.superClass)
}

// ConstantPool
func (self *ClassFile) ConstantPool() ConstantPool {
	return self.constantPool
}

// Getter InterfaceNames
func (self *ClassFile) InterfaceNames() []uint16 {
	return self.interfaces
}
