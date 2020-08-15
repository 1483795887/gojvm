package math

import (
	"gojvm/ch05/instructions/base"
	"gojvm/ch05/rtda"
)

//IOr int或运算
type IOr struct {
	base.NoOperandsInstruction
}

func (client *IOr) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v2 | v1
	stack.PushInt(result)
}

//LOr long或运算
type LOr struct {
	base.NoOperandsInstruction
}

func (client *LOr) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v2 | v1
	stack.PushLong(result)
}
