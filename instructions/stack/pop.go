package stack

import (
	"gojvm/ch05/instructions/base"
	"gojvm/ch05/rtda"
)

//Pop 弹出一个
type Pop struct {
	base.NoOperandsInstruction
}

func (client *Pop) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

//Pop2 弹出两个
type Pop2 struct {
	base.NoOperandsInstruction
}

func (client *Pop2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
