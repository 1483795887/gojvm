package math

import (
	"gojvm/ch05/instructions/base"
	"gojvm/ch05/rtda"
)

//IInc 自增
type IInc struct {
	Index uint
	Const int32
}

func (client *IInc) FetchOperands(reader *base.BytecodeReader) {
	client.Index = uint(reader.ReadUint8())
	client.Const = int32(reader.ReadInt8())
}

func (client *IInc) Execute(frame *rtda.Frame) {
	localVars := frame.LocalVars()
	val := localVars.GetInt(client.Index)
	val += client.Const
	localVars.SetInt(client.Index, val)
}
