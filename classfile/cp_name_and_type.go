package classfile

//ConstantNameAndTypeInfo 表示名字和类型
type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (client *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	client.nameIndex = reader.readUint16()
	client.descriptorIndex = reader.readUint16()
}
