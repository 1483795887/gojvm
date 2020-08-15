package classfile

import "fmt"

//ClassFile 描述java class文件内容
type ClassFile struct {
	magic        uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

//Parse 解析文件
func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
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

func (client *ClassFile) read(reader *ClassReader) {
	client.readAndCheckMagic(reader)
	client.readAndCheckVersion(reader)

	client.constantPool = readConstantPool(reader)

	client.accessFlags = reader.readUint16()
	client.thisClass = reader.readUint16()
	client.superClass = reader.readUint16()
	client.interfaces = reader.readUint16s()

	client.fields = readMembers(reader, client.constantPool)

	client.methods = readMembers(reader, client.constantPool)
	client.attributes = readAttributes(reader, client.constantPool)
}

func (client *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xcafebabe {
		panic("java.lang.ClassFormatError: magic!")
	}
}

func (client *ClassFile) readAndCheckVersion(reader *ClassReader) {
	client.minorVersion = reader.readUint16()
	client.majorVersion = reader.readUint16()
	if client.majorVersion == 45 {
		return
	}
	if client.majorVersion > 45 && client.minorVersion == 0 {
		return
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

//MajorVersion 获得主版本
func (client *ClassFile) MajorVersion() uint16 {
	return client.majorVersion
}

//MinorVersion 获得次版本
func (client *ClassFile) MinorVersion() uint16 {
	return client.minorVersion
}

//ClassName 类名
func (client *ClassFile) ClassName() string {
	return client.constantPool.getClassName(client.thisClass)
}

//SuperClassName 父类名
func (client *ClassFile) SuperClassName() string {
	if client.superClass > 0 {
		return client.constantPool.getClassName(client.superClass)
	}
	return ""
}

//InterfaceNames 查找接口名
func (client *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(client.interfaces))
	for i, cpIndex := range client.interfaces {
		interfaceNames[i] = client.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}

//ConstantPool 返回所有常量
func (client *ClassFile) ConstantPool() []ConstantInfo{
	return client.constantPool
}

//AccessFlag 返回访问标志
func (client *ClassFile) AccessFlag() uint16{
	return client.accessFlags
}

//Fields 返回成员
func (client *ClassFile) Fields() []*MemberInfo{
	return client.fields
}

//Methods 方法数
func (client *ClassFile) Methods() []*MemberInfo{
	return client.methods
}