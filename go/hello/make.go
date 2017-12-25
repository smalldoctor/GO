package main

import "fmt"

/*
make(T, args) 返回的是 T 的 引用
如果不特殊声明，go 的函数默认都是按值穿参，即通过函数传递的参数是值的副本，
在函数内部对值修改不影响值的本身，但是 make(T, args) 返回的值通过函数传递参数之后可以直接修改，
即 map，slice，channel 通过函数穿参之后在函数内部修改将影响函数外部的值。
*/
func main() {
	var s1 []int
	if s1 == nil {
		fmt.Printf("s1 is nil --> %#v \n ", s1) // []int(nil)
	}
	s2 := make([]int, 3)
	if s2 == nil {
		fmt.Printf("s2 is nil --> %#v \n ", s2)
	} else {
		fmt.Printf("s2 is not nill --> %#v \n ", s2) // []int{0, 0, 0}
	}
	var m1 map[int]string
	if m1 == nil {
		fmt.Printf("m1 is nil --> %#v \n ", m1) //map[int]string(nil)
	}
	m2 := make(map[int]string)
	if m2 == nil {
		fmt.Printf("m2 is nil --> %#v \n ", m2)
	} else {
		fmt.Printf("m2 is not nill --> %#v \n ", m2) //map[int]string{}
	}
	var c1 chan string
	if c1 == nil {
		fmt.Printf("c1 is nil --> %#v \n ", c1) //(chan string)(nil)
	}
	c2 := make(chan string)
	if c2 == nil {
		fmt.Printf("c2 is nil --> %#v \n ", c2)
	} else {
		fmt.Printf("c2 is not nill --> %#v \n ", c2) //(chan string)(0xc420016120)
	}
}
