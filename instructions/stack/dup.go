package stack

import (
	"gojvm/ch05/instructions/base"
	"gojvm/ch05/rtda"
)

//Dup A->A,A
type Dup struct {
	base.NoOperandsInstruction
}

func (client *Dup) Execute(frame *rtda.Frame) {
	slot := frame.OperandStack().PopSlot()
	stack := frame.OperandStack()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}

//DupX1 A,B->A,B,A
type DupX1 struct {
	base.NoOperandsInstruction
}

func (client *DupX1) Execute(frame *rtda.Frame) {
	slotA := frame.OperandStack().PopSlot()
	slotB := frame.OperandStack().PopSlot()
	stack := frame.OperandStack()
	stack.PushSlot(slotA)
	stack.PushSlot(slotB)
	stack.PushSlot(slotA)
}

//DupX2 A,B,C->A,B,C,A
type DupX2 struct {
	base.NoOperandsInstruction
}

func (client *DupX2) Execute(frame *rtda.Frame) {
	slotA := frame.OperandStack().PopSlot()
	slotB := frame.OperandStack().PopSlot()
	slotC := frame.OperandStack().PopSlot()
	stack := frame.OperandStack()
	stack.PushSlot(slotA)
	stack.PushSlot(slotC)
	stack.PushSlot(slotB)
	stack.PushSlot(slotA)
}

//Dup2 A,B->A,B,A,B
type Dup2 struct {
	base.NoOperandsInstruction
}

func (client *Dup2) Execute(frame *rtda.Frame) {
	slotA := frame.OperandStack().PopSlot()
	slotB := frame.OperandStack().PopSlot()
	stack := frame.OperandStack()
	stack.PushSlot(slotB)
	stack.PushSlot(slotA)
	stack.PushSlot(slotB)
	stack.PushSlot(slotA)
}

//Dup2X1 A,B,C->A,B,C,A,B
type Dup2X1 struct {
	base.NoOperandsInstruction
}

func (client *Dup2X1) Execute(frame *rtda.Frame) {
	slotA := frame.OperandStack().PopSlot()
	slotB := frame.OperandStack().PopSlot()
	slotC := frame.OperandStack().PopSlot()
	stack := frame.OperandStack()
	stack.PushSlot(slotB)
	stack.PushSlot(slotA)
	stack.PushSlot(slotC)
	stack.PushSlot(slotB)
	stack.PushSlot(slotA)
}

//Dup2X2 A,B,C,D->A,B,C,D,A,B
type Dup2X2 struct {
	base.NoOperandsInstruction
}

func (client *Dup2X2) Execute(frame *rtda.Frame) {
	slotA := frame.OperandStack().PopSlot()
	slotB := frame.OperandStack().PopSlot()
	slotC := frame.OperandStack().PopSlot()
	slotD := frame.OperandStack().PopSlot()
	stack := frame.OperandStack()
	stack.PushSlot(slotB)
	stack.PushSlot(slotA)
	stack.PushSlot(slotD)
	stack.PushSlot(slotC)
	stack.PushSlot(slotB)
	stack.PushSlot(slotA)
}
