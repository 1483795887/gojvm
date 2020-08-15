package base

import (
	"gojvm/ch05/rtda"
)

//StoreOperation store接口
type StoreOperation interface {
	Store(frame *rtda.Frame, index uint)
}

//IStoreOperation 加载int
type IStoreOperation struct{}

func (client *IStoreOperation) Store(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopInt()
	frame.LocalVars().SetInt(index, val)
}

//LStoreOperation 加载long
type LStoreOperation struct{}

func (client *LStoreOperation) Store(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}

//FStoreOperation 加载float
type FStoreOperation struct{}

func (client *FStoreOperation) Store(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopFloat()
	frame.LocalVars().SetFloat(index, val)
}

//DStoreOperation 加载double
type DStoreOperation struct{}

func (client *DStoreOperation) Store(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(index, val)
}

//AStoreOperation 加载double
type AStoreOperation struct{}

func (client *AStoreOperation) Store(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopRef()
	frame.LocalVars().SetRef(index, val)
}

//Store 带操作数的存储指令
type Store struct {
	Index8Instruction
	op StoreOperation
}

func (client *Store) Execute(frame *rtda.Frame) {
	client.op.Store(frame, uint(client.Index))
}

//StoreN 不带操作数的存储指令
type StoreN struct {
	NoOperandsInstruction
	op    StoreOperation
	Index uint
}

func (client *StoreN) Execute(frame *rtda.Frame) {
	client.op.Store(frame, uint(client.Index))
}
