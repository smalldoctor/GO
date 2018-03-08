package rtda

/*
局部变量表的元素,每个元素需要能存放一个int或者指针；
两个局部变量元素标示一个long或者double
*/

type Slot struct {
	num int
	ref *Object
}
