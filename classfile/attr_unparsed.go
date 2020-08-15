package classfile

//UnparsedAttribute 未解析的属性
type UnparsedAttribute struct{
	name string
	length uint32
	info []byte
}

func (client *UnparsedAttribute) readInfo(reader * ClassReader){
	client.info = reader.readBytes(client.length)
}