package base

import "gojvm/ch05/rtda"

//Instruction 指令接口
type Instruction interface {
	FetchOperands(reader *BytecodeReader)
	Execute(frame *rtda.Frame)
}

//NoOperandsInstruction 无操作数指令
type NoOperandsInstruction struct{}

func (client *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// 什么都不做
}

//BranchInstruction 跳转指令
type BranchInstruction struct {
	offset int
}

func (client *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	client.offset = int(reader.ReadInt16())
}

//Index8Instruction 8位引索指令
type Index8Instruction struct {
	Index int
}

func (client *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	client.Index = int(reader.ReadInt8())
}

//Index16Instruction 16位引索指令
type Index16Instruction struct {
	Index int
}

func (client *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	client.Index = int(reader.ReadInt16())
}
