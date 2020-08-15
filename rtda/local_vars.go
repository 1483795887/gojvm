package rtda

import "math"

//LocalVars 表示局部变量表
type LocalVars []Slot

func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}

//SetInt 设置一个int类型变量
func (client LocalVars) SetInt(index uint, val int32) {
	client[index].num = val
}

//GetInt 获取一个int类型变量
func (client LocalVars) GetInt(index uint) int32 {
	return client[index].num
}

//SetFloat 设置一个flaot类型变量
func (client LocalVars) SetFloat(index uint, val float32) {
	client[index].num = int32(math.Float32bits(val))
}

//GetFloat 设置一个flaot类型变量
func (client LocalVars) GetFloat(index uint) float32 {
	bits := uint32(client[index].num)
	return math.Float32frombits(bits)
}

//SetLong 设置一个long类型变量
func (client LocalVars) SetLong(index uint, val int64) {
	client[index].num = int32(val)
	client[index+1].num = int32(val >> 32)
}

//GetLong 获取一个long类型变量
func (client LocalVars) GetLong(index uint) int64 {
	low := uint32(client[index].num)
	high := uint32(client[index+1].num)
	return int64(high)<<32 | int64(low)
}

//SetDouble 设置一个double类型变量
func (client LocalVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	client.SetLong(index, int64(bits))
}

//GetDouble 设置一个double类型变量
func (client LocalVars) GetDouble(index uint) float64 {
	bits := uint64(client.GetLong(index))
	return math.Float64frombits(bits)
}

//SetRef 设置一个ref类型变量
func (client LocalVars) SetRef(index uint, ref *Object) {
	client[index].ref = ref
}

//GetRef 获取一个ref类型变量
func (client LocalVars) GetRef(index uint) *Object {
	return client[index].ref
}