package classfile

import "math"

//ConstantIntegerInfo 整形常量
type ConstantIntegerInfo struct {
	val int32
}

func (client *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	client.val = int32(bytes)
}

//ConstantFloatInfo 浮点常量
type ConstantFloatInfo struct {
	val float32
}

func (client *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	client.val = math.Float32frombits(bytes)
}

//ConstantLongInfo 长整型常量
type ConstantLongInfo struct {
	val uint64
}

func (client *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	client.val = uint64(bytes)
}

//ConstantDoubleInfo 双精度常量
type ConstantDoubleInfo struct {
	val float64
}

func (client *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	client.val = math.Float64frombits(bytes)
}