package constants

import (
	"gojvm/ch05/instructions/base"
	"gojvm/ch05/rtda"
)

//BiPush 压入byte
type BiPush struct {
	val int8
}

func (client *BiPush) FetchOperands(reader *base.BytecodeReader) {
	client.val = reader.ReadInt8()
}

func (client *BiPush) Execute(frame *rtda.Frame) {
	i := int32(client.val)
	frame.OperandStack().PushInt(i)
}

//SiPush 压入short
type SiPush struct {
	val int16
}

func (client *SiPush) FetchOperands(reader *base.BytecodeReader) {
	client.val = reader.ReadInt16()
}

func (client *SiPush) execute(frame *rtda.Frame) {
	i := int32(client.val)
	frame.OperandStack().PushInt(i)
}
