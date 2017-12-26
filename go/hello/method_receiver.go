package main

type Fooer interface {
	Foo()
}

type A struct{}
type B struct{}

func (a A) Foo()  {}
func (b *B) Foo() {}

func main() {
	var v Fooer
	var a A
	var b B
	/*
	在理解方法的接收者是指针类型和非指针类型之前，请先理解go中的变量：
	在go中，变量是具有类型（与python不一样），指针是变量的类型一种；
	所以方法的接受者是在指定什么样的变量具有方法。如(a A) Foo() 是在
	指定结构体A类型的变量具有方法，（在调用这个方法时，其实是将变量的
	值传递到方法中去）；(b *B) Foo() 是B类型结构体的指针具有这个
	方法（在调用这个方法时，其实是将B类型结构体指针变量传递到方法中）

	对于接口，go认为只要是实现接口的所有方法的变量，就是变量的实现；
	所以在例子中(b *B) Foo() 是B类型结构体指针变量实现了方法，所以
	必须是B类型结构体指针变量，而不是B类型结构体
	*/
	v = a // This is OK, A has a method Foo()
	//v = b  // This is *NOT* OK. B *DOES NOT* have a method Foo().
	v = &b // This is OK. *B have a method Foo()
	_ = v
}
