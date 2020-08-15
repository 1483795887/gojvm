package rtda

//Slot 表示局部变量，num表示数值变量，ref表示引用变量
type Slot struct {
	num int32
	ref *Object
}