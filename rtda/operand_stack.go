package rtda

import "math"

//OperandStack 操作数栈
type OperandStack struct {
	size  uint
	slots []Slot
}

func newOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{
			slots: make([]Slot, maxStack),
		}
	}
	return nil
}

//PushInt 压入int型变量
func (client *OperandStack) PushInt(val int32) {
	client.slots[client.size].num = val
	client.size++
}

//PopInt 弹出int型变量
func (client *OperandStack) PopInt() int32 {
	client.size--
	return client.slots[client.size].num
}

//PushFloat 压入float型变量
func (client *OperandStack) PushFloat(val float32) {
	client.slots[client.size].num = int32(math.Float32bits(val))
	client.size++
}

//PopFloat 弹出float型变量
func (client *OperandStack) PopFloat() float32 {
	client.size--
	bits := uint32(client.slots[client.size].num)
	return math.Float32frombits(bits)
}

//PushLong 压入long型变量
func (client *OperandStack) PushLong(val int64) {
	client.slots[client.size].num = int32(val)
	client.slots[client.size+1].num = int32(val >> 32)
	client.size += 2
}

//PopLong 弹出long型变量
func (client *OperandStack) PopLong() int64 {
	client.size -= 2
	low := uint32(client.slots[client.size].num)
	high := uint32(client.slots[client.size+1].num)
	return int64(high)<<32 | int64(low)
}

//PushDouble 压入double型变量
func (client *OperandStack) PushDouble(val float64) {
	bits := math.Float64bits(val)
	client.PushLong(int64(bits))
}

//PopDouble 弹出double型变量
func (client *OperandStack) PopDouble() float64 {
	bits := uint64(client.PopLong())
	return math.Float64frombits(bits)
}

//PushRef 压入ref型变量
func (client *OperandStack) PushRef(ref *Object) {
	client.slots[client.size].ref = ref
	client.size++
}

//PopRef 弹出ref型变量
func (client *OperandStack) PopRef() *Object {
	client.size--
	return client.slots[client.size].ref
}

//PushSlot 压入与变量类型无关的slot
func (client *OperandStack) PushSlot(slot Slot) {
	client.slots[client.size] = slot
	client.size++
}

//PopSlot 弹出与变量类型无关的slot
func (client *OperandStack) PopSlot() Slot {
	client.size--
	return client.slots[client.size]
}
