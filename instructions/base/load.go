package base

import (
	"gojvm/ch05/rtda"
)

//LoadOperation load接口
type LoadOperation interface {
	Load(frame *rtda.Frame, index uint)
}

//ILoadOperation 加载int
type ILoadOperation struct{}

func (client *ILoadOperation) Load(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}

//LLoadOperation 加载long
type LLoadOperation struct{}

func (client *LLoadOperation) Load(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetLong(index)
	frame.OperandStack().PushLong(val)
}

//FLoadOperation 加载float
type FLoadOperation struct{}

func (client *FLoadOperation) Load(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetFloat(index)
	frame.OperandStack().PushFloat(val)
}

//DLoadOperation 加载double
type DLoadOperation struct{}

func (client *DLoadOperation) Load(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetDouble(index)
	frame.OperandStack().PushDouble(val)
}

//ALoadOperation 加载double
type ALoadOperation struct{}

func (client *ALoadOperation) Load(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetRef(index)
	frame.OperandStack().PushRef(val)
}

//Load 带操作数的加载指令
type Load struct {
	Index8Instruction
	op LoadOperation
}

func (client *Load) Execute(frame *rtda.Frame) {
	client.op.Load(frame, uint(client.Index))
}

//LoadN 不带操作数的加载指令
type LoadN struct {
	NoOperandsInstruction
	op    LoadOperation
	Index uint
}

func (client *LoadN) Execute(frame *rtda.Frame) {
	client.op.Load(frame, uint(client.Index))
}
