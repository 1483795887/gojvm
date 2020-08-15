package constants

import (
	"gojvm/ch05/instructions/base"
	"gojvm/ch05/rtda"
)

//ConstRefNull 加载常量空指针
type ConstRefNull struct{
	base.NoOperandsInstruction
}

func (client *ConstRefNull) Execute(frame *rtda.Frame){
	frame.OperandStack().PushRef(nil)
}

//ConstDouble 加载常量double
type ConstDouble struct{
	base.NoOperandsInstruction
	val float64
}

func (client *ConstDouble) Execute(frame *rtda.Frame){
	frame.OperandStack().PushDouble(client.val)
}

//ConstFloat 加载常量float
type ConstFloat struct{
	base.NoOperandsInstruction
	val float32
}

func (client *ConstFloat) Execute(frame *rtda.Frame){
	frame.OperandStack().PushFloat(client.val)
}

//ConstLong 加载常量long
type ConstLong struct{
	base.NoOperandsInstruction
	val int64
}

func (client *ConstLong) Execute(frame *rtda.Frame){
	frame.OperandStack().PushLong(client.val)
}

//ConstInt 加载常量int
type ConstInt struct{
	base.NoOperandsInstruction
	val int32
}

func (client *ConstInt) Execute(frame *rtda.Frame){
	frame.OperandStack().PushInt(client.val)
}