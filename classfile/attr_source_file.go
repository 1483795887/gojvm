package classfile

//SourceFileAttribute 源文件属性
type SourceFileAttribute struct{
	cp ConstantPool
	sourceFileIndex uint16
}

func (client *SourceFileAttribute) readInfo(reader *ClassReader){
	client.sourceFileIndex = reader.readUint16()
}

//FileName 返回源文件的文件名
func (client *SourceFileAttribute) FileName() string{
	return client.cp.getUtf8(client.sourceFileIndex)
}