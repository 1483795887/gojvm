package classfile

//ExceptionTableEntry 异常处理表项
type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

//CodeAttribute 代码属性
type CodeAttribute struct {
	cp             ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributes     []AttributeInfo
}

func (client *CodeAttribute) readInfo(reader *ClassReader) {
	client.maxLocals = reader.readUint16()
	client.maxLocals = reader.readUint16()
	codeLength := reader.readUint32()
	client.code = reader.readBytes(codeLength)
	client.exceptionTable = readExceptionTable(reader)
	client.attributes = readAttributes(reader, client.cp)
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := reader.readUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc:   reader.readUint16(),
			endPc:     reader.readUint16(),
			handlerPc: reader.readUint16(),
			catchType: reader.readUint16(),
		}
	}
	return exceptionTable
}