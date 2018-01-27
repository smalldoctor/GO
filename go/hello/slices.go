package main

import (
	"fmt"
	"reflect"
)

/*
切片是对一个已经存在数组的部分引用，相当于定一个了指针；
var x = []int{2, 3, 5, 7, 11} =>定义了一个数组，同时x是指向这个数组的切片；
切片在内存中的组织方式实际上是一个有 3 个域的结构体：指向相关数组的指针，切片长度以及切片容量；
*/
// 可以接受一个切片
func sum(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}

func main() {
	var arr1 [6]int
	// 切片的类型还是数组
	var slice1 []int = arr1[2:5] // item at index 5 not included!

	// load the array with integers: 0,1,2,3,4,5
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	// print the slice
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	fmt.Printf("The length of arr1 is %d\n", len(arr1))
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	// grow the slice
	slice1 = slice1[0:4]
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	sum(arr1[:])

	slice1 = []int{}
	fmt.Println(len(slice1)) // 0

	// append函数在slice中追加元素，如已经没有容量了，则会新生成一个数组，
	arr3 := [1]int{1}
	slice3 := arr3[:]
	fmt.Println("arr3: ", arr3, " slice3: ", slice3) // arr3:  [1]  slice3:  [1]
	slice3 = append(slice3, 2, 3)
	fmt.Println("arr3: ", arr3, " slice3: ", slice3) // arr3:  [1]  slice3:  [1 2 3]

	// append函数在slice中追加元素时，如果存在容量，则会在扩大长度，追加内容；因为slice是对数组的引用，所以修改数组对应位置的内容
	arr4 := [5]int{0, 1}
	slice4 := arr4[1:3]
	fmt.Println("arr4: ", arr4, " slice4: ", slice4) //arr4:  [0 1 0 0 0]  slice4:  [1 0]
	slice4 = append(slice4, 5, 6)
	fmt.Println("arr4: ", arr4, " slice4: ", slice4) //arr4:  [0 1 0 5 6]  slice4:  [1 0 5 6]
	// grow the slice beyond capacity
	//slice1 = slice1[0:7 ] // panic: runtime error: slice bound out of range

	slicePoint := &slice4
	// 切片类型是数组，只不过是他本身所引用的数组共享了内存空间；切片在实现上使用了自己的数据模型，
	// 从而可以进行比普通数组更多的操作；可以类比JAVA的集合类型，切片底层是用数组实现的；
	fmt.Println("slicePoint: ", slicePoint)
	fmt.Println("the array of slice4: ", &arr4)

	var slicePoint2 *[]int
	slicePoint2 = &slice4
	fmt.Println("slicePoint2: ", slicePoint2)
	fmt.Println("slicePoint2 type: ", reflect.TypeOf(slicePoint2), " slice4 type: ", reflect.TypeOf(slice4))

}
