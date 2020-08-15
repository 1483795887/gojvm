package classfile

//ConstantClassInfo 表示class
type ConstantClassInfo struct{
	cp ConstantPool
	nameIndex uint16
}

func (client *ConstantClassInfo) readInfo(reader *ClassReader){
	client.nameIndex = reader.readUint16()
}

//Name 返回类名
func (client *ConstantClassInfo) Name() string{
	return client.cp.getUtf8(client.nameIndex)
}