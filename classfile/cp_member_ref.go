package classfile

//ConstantMemberrefInfo 表示成员参考
type ConstantMemberrefInfo struct {
	cp ConstantPool
	classIndex uint16
	nameAndTypeIndex uint16
}

//ConstantFieldrefInfo 成员参考
type ConstantFieldrefInfo struct{ConstantMemberrefInfo}
//ConstantMethodrefInfo 方法参考
type ConstantMethodrefInfo struct {ConstantMemberrefInfo}
//ConstantInterfaceMethodrefInfo 接口参考
type ConstantInterfaceMethodrefInfo struct{ConstantMemberrefInfo}

func (client *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	client.classIndex = reader.readUint16()
	client.nameAndTypeIndex = reader.readUint16()
}

//ClassName 返回类名
func (client *ConstantMemberrefInfo) ClassName() string{
	return client.cp.getClassName(client.classIndex)
}

//NameAndDescriptor 返回类型和描述符
func (client *ConstantMemberrefInfo) NameAndDescriptor() (string, string){
	return client.cp.getNameAndType(client.nameAndTypeIndex)
}