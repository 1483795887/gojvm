package rtda

//Stack 表示虚拟机栈
type Stack struct {
	maxSize uint
	size    uint
	_top    *Frame
}

func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

func (client *Stack) push(frame *Frame) {
	if client.size >= client.maxSize {
		panic("java.lang.StackOverflowError")
	}
	if client._top != nil {
		frame.lower = client._top
	}
	client._top = frame
	client.size++
}

func (client *Stack) pop() *Frame {
	if client._top == nil {
		panic("jvm stack is empty!")
	}
	top := client._top
	client._top = top.lower
	top.lower = nil
	client.size--
	return top
}

func (client *Stack) top() *Frame {
	if client._top == nil {
		panic("jvm stack is empty!")
	}
	return client._top
}
