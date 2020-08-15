package stack

import (
	"gojvm/ch05/instructions/base"
	"gojvm/ch05/rtda"
)

//Swap A,B->B,A
type Swap struct {
	base.NoOperandsInstruction
}

func (client *Swap) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}
