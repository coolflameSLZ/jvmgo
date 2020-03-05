package rtda

/**
num字段存放整数，ref字段存放引用
*/
type Slot struct {
	num int32
	ref *Object
}
