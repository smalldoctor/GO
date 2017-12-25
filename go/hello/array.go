package main

import "fmt"

/*
arr1是一个指针变量,一个指向数组的指针；
arr2是一个普通整型数组;
*/
var arr1 = new([5]int)
// 值类型
var arr2 [5]int

// 数组的值传递，会拷贝数组，性能差
func f(a [3]int)  {

}
// 数组的指针传递，性能更好
func fp(a *[3]int)  {

}

func main() {
	// 是一个拷贝的过程，数组类型是值类型
	arr2 := *arr1

	//对指定位置进行初始化
	var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}
	fmt.Println(arrKeyValue)
	fmt.Println(arr1)
	fmt.Println(arr2)

	fmt.Println(arr1[3])
	fmt.Println(&arr1[3])

	arr3 := new([]int)
	fmt.Println(arr3)
	fmt.Println(*arr3)

	x := new(string)
	fmt.Println(x)
	// 默认对应类型的零值；string的零值是空字符串
	fmt.Println(*x)
}
