package classfile

import "encoding/binary"

//ClassReader 负责读取class
type ClassReader struct {
	data []byte
}

func (client *ClassReader) readUint8() uint8 {
	val := client.data[0]
	client.data = client.data[1:]
	return val
}

func (client *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(client.data)
	client.data = client.data[2:]
	return val
}

func (client *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(client.data)
	client.data = client.data[4:]
	return val
}

func (client *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(client.data)
	client.data = client.data[8:]
	return val
}

func (client *ClassReader) readUint16s() []uint16 {
	n := client.readUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = client.readUint16()
	}
	return s
}

func (client *ClassReader) readBytes(n uint32) []byte {
	bytes := client.data[:n]
	client.data = client.data[n:]
	return bytes
}
