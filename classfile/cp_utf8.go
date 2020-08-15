package classfile

//ConstantUtf8Info 是utf8对应的常量
type ConstantUtf8Info struct {
	str string
}

func (client *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	client.str = decodeMUTF8(bytes)
}

func decodeMUTF8(bytes []byte) string {
	return string(bytes)	//要求不出现增补字符和null字符
}
