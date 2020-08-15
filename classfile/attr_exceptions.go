package classfile

//ExceptionAttribute 记录方法抛出的异常表
type ExceptionAttribute struct {
	exceptionIndexTable []uint16
}

func (client *ExceptionAttribute) readInfo(reader *ClassReader) {
	client.exceptionIndexTable = reader.readUint16s()
}

//ExceptionIndexTable 返回异常表
func (client *ExceptionAttribute) ExceptionIndexTable() []uint16 {
	return client.exceptionIndexTable
}
