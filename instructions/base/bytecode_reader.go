package base

//BytecodeReader 字节码读取器
type BytecodeReader struct {
	code []byte
	pc   int
}

//Reset 重置
func (client *BytecodeReader) Reset(code []byte, pc int) {
	client.code = code
	client.pc = pc
}

//ReadUint8 读取一个uint8
func (client *BytecodeReader) ReadUint8() uint8 {
	i := client.code[client.pc]
	client.pc++
	return i
}

//ReadInt8 读取一个int8
func (client *BytecodeReader) ReadInt8() int8 {
	return int8(client.ReadUint8())
}

//ReadUint16 读取一个uint16
func (client *BytecodeReader) ReadUint16() uint16 {
	byte1 := uint16(client.ReadUint8())
	byte2 := uint16(client.ReadUint8())
	return (byte1 << 8) | byte2
}

//ReadInt16 读取一个int16
func (client *BytecodeReader) ReadInt16() int16 {
	return int16(client.ReadUint16())
}

//ReadInt32 读取一个uint16
func (client *BytecodeReader) ReadInt32() int32 {
	byte1 := int32(client.ReadUint8())
	byte2 := int32(client.ReadUint8())
	byte3 := int32(client.ReadUint8())
	byte4 := int32(client.ReadUint8())
	client.pc += 4
	return (byte1 << 24) | (byte2 << 16) | (byte3 << 8) | byte4
}
