package rtda

//Thread 表示java中的线程
type Thread struct {
	pc    int
	stack *Stack
}

//NewThread 创建线程
func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}


//PushFrame 压入栈帧
func (client *Thread) PushFrame(frame *Frame){
	client.stack.push(frame)
}

//CurrentFrame 返回当前栈帧
func (client *Thread) CurrentFrame() *Frame{
	return client.stack.top()
}