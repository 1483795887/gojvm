package math

import (
	"gojvm/ch05/instructions/base"
	"gojvm/ch05/rtda"
)

//IShl int左移
type IShl struct {
	base.NoOperandsInstruction
}

func (client *IShl) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := v1 << s
	stack.PushInt(result)
}

//IShr int算术右移
type IShr struct {
	base.NoOperandsInstruction
}

func (client *IShr) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := v1 >> s
	stack.PushInt(result)
}

//IUShr int逻辑右移
type IUShr struct {
	base.NoOperandsInstruction
}

func (client *IUShr) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := int32(uint32(v1) >> s)
	stack.PushInt(result)
}

//LShl long左移
type LShl struct {
	base.NoOperandsInstruction
}

func (client *LShl) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := v1 << s
	stack.PushLong(result)
}

//LShr long算术右移
type LShr struct {
	base.NoOperandsInstruction
}

func (client *LShr) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := v1 >> s
	stack.PushLong(result)
}

//LUShr long逻辑右移
type LUShr struct {
	base.NoOperandsInstruction
}

func (client *LUShr) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := int64(uint64(v1) >> s)
	stack.PushLong(result)
}
