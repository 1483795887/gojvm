package math

import (
	"gojvm/ch05/instructions/base"
	"gojvm/ch05/rtda"
	"math"
)

//IRem int取余
type IRem struct {
	base.NoOperandsInstruction
}

func (client *IRem) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	result := v1 % v2
	stack.PushInt(result)
}

//LRem long取余
type LRem struct {
	base.NoOperandsInstruction
}

func (client *LRem) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	result := v1 % v2
	stack.PushLong(result)
}

//FRem float取余
type FRem struct {
	base.NoOperandsInstruction
}

func (client *FRem) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := float32(math.Mod(float64(v1), float64(v2)))
	stack.PushFloat(result)
}

//DRem floag取余
type DRem struct {
	base.NoOperandsInstruction
}

func (client *DRem) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := math.Mod(v1, v2)
	stack.PushDouble(result)
}
