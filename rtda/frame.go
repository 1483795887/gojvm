package rtda

//Frame 表示栈帧
type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
}

//NewFrame 新建一个帧
func NewFrame(maxLocals, maxStacks uint) *Frame {
	return &Frame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStacks),
	}
}

//LocalVars 获取局部变量表
func (client *Frame) LocalVars() LocalVars {
	return client.localVars
}

//OperandStack 获取操作数栈
func (client *Frame) OperandStack() *OperandStack {
	return client.operandStack
}
