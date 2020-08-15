package classfile

//ConstantValueAttribute 常量属性
type ConstantValueAttribute struct{
	constantValueIndex uint16
}

func (client *ConstantValueAttribute) readInfo(reader *ClassReader){
	client.constantValueIndex = reader.readUint16()
}

//ConstantValueIndex 返回常数属性
func (client *ConstantValueAttribute) ConstantValueIndex() uint16{
	return client.constantValueIndex
}