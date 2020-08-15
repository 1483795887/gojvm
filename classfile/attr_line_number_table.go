package classfile

//LineNumberTableAttribute 记录方法抛出的异常表
type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}

//LineNumberTableEntry 行号
type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (client *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	lineNumberTableLength := reader.readUint16()
	client.lineNumberTable = make([]*LineNumberTableEntry, lineNumberTableLength)
	for i := range client.lineNumberTable {
		client.lineNumberTable[i] = &LineNumberTableEntry{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}
