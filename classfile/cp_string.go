package classfile


//ConstantStringInfo 对应String类型的常量
type ConstantStringInfo struct{
	cp ConstantPool
	stringIndex uint16
}

func (client *ConstantStringInfo) readInfo(reader *ClassReader){
	client.stringIndex = reader.readUint16()
}

func (client *ConstantStringInfo) String() string{
	return client.cp.getUtf8(client.stringIndex)
}