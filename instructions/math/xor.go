package math

import (
	"gojvm/ch05/instructions/base"
	"gojvm/ch05/rtda"
)

//IXor int异或运算
type IXor struct {
	base.NoOperandsInstruction
}

func (client *IXor) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v2 ^ v1
	stack.PushInt(result)
}

//LXor long异或运算
type LXor struct {
	base.NoOperandsInstruction
}

func (client *LXor) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v2 ^ v1
	stack.PushLong(result)
}
