package constants

import (
	"gojvm/ch05/instructions/base"
	"gojvm/ch05/rtda"
)

//NOP 空指令
type NOP struct {
	base.NoOperandsInstruction
}

func (client *NOP) execute(frame *rtda.Frame) {
	// 什么都不做
}
