package math

import (
	"gojvm/ch05/instructions/base"
	"gojvm/ch05/rtda"
)

//IAnd int与运算
type IAnd struct {
	base.NoOperandsInstruction
}

func (client *IAnd) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v2 & v1
	stack.PushInt(result)
}

//LAnd long与运算
type LAnd struct {
	base.NoOperandsInstruction
}

func (client *LAnd) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v2 & v1
	stack.PushLong(result)
}
