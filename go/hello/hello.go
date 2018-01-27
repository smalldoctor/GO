/*
每个文件都必须指定其所属于的包；每个go程序都需要有一个main包；
包中的内容被导出之后，则可以在包外被使用;
导出：通过标识符的首字母大写的方式实现导出；
GO访问控制只分两种，首字母大写的变量，函数，结构体，类型等可以在包外访问；
如果首字母小写的变量，函数，结构体，类型等不可以在包外访问,只能包内访问;
 */
package main

/*
import 后面的是包的路径，而非包的名称；只是在通常情况下包名和路径名相同；
如 fmt 是在GO安装目录的src下的fmt的目录
*/
import (
	"fmt"
	"go/stringutil"
	"math"
	"os"
	"strconv"
)

var (
	// 全局变量,全局变量使用这种方式;
	// GO中的函数内定义的变量为局部变量；函数外的为全局变量
	// 局部变量与全局变量同名时，优先使用局部变量

	g      int
	slice1 = []int{1, 2}
)

/**
	* 定义常量
	* 常量可以使用内置函数计算的值
 */
const c1 = 1
const c2 = 2
const (
	c3 = 3
)

func test() {
	var (
		sdf int
	)
	fmt.Print(sdf)
}

/**
	* 每个GO程序都必须有一个main函数；首先执行init函数，然后执行main函数
 */
func main() {

	go test()
	/**
		* GO语言每行代表语句的结束；如果多个语句在同一行，必须以;进行分隔
	 */
	fmt.Printf(stringutil.Reverse("!oG ,olleH"))

	// 指定类型并初始化
	var a int = 10
	// 根据初始化的值自动识别类型
	var b = "b"
	// 只能在函数体内使用这种方式
	c := 1.2
	// a := 1.3 error；因为a是已经存在的变量，即:=不能用于已经存在的变量
	// a = "error";c = "error" error;类型确定之后不能再赋值不兼容的类型
	fmt.Println(a, b, c)
	// _, y = 5, 7 5会被舍弃，_是一个只写变量，不能被读取；可以用于对返回多个返回值的函数的返回值进行部分舍弃
	// _ 如果用在import 中表示只调用相关包的init函数，但是不能通过包名访问其中的内容

	/**
		* 函数作为值
		* 函数作为闭包，闭包是匿名函数，是函数内的局部变量不被释放，并且达到封装的效果
		* 有接收者的函数被作为方法,
	 */
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	fmt.Println(getSquareRoot(9))

	//方法
	var c1, c2, c3, c4 Circle
	c1.radius = 8
	fmt.Println("Area of Circle(c1) = ", c1.getArea())
	fmt.Println("radius of circle = ", c1.radius)
	c2.radius = 8
	fmt.Println("Area of Circle(c1) = ", c2.getArea2())
	fmt.Println("radius of circle = ", c2.radius)

	/**
		* 测试指针传递
		* GO中的函数默认是值传递
	 */
	c3.radius = 8
	// 值传递
	testRef(c3)
	fmt.Println("radius of circle = ", c3.radius)
	c4.radius = 8
	// 引用传递
	testRef2(&c4)
	fmt.Println("radius of circle = ", c4.radius)

	/**
		* range用来迭代数组，切片，返回索引和元素
		* range用来迭代map，返回key和value
		* range用来迭代Unicode字符串，返回字符索引和Unicode值
	 */

	/**
		* map[key_type] value_type或者使用make函数声明map
		* delete(key)用于删除map的key
	 */

	fileInfo, err := os.Stat("/Library/Java/JavaVirtualMachines/1.6.0.jdk/Contents/Home")
	if err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Println("isDir: " + strconv.FormatBool(fileInfo.IsDir()))
	}

}

type Circle struct {
	radius float64
}

//该 method 属于 Circle 类型对象中的方法;
//相当于值传递
func (c Circle) getArea() float64 {
	//c.radius 即为 Circle 类型对象中的属性
	var area = 3.14 * c.radius * c.radius
	// 不会改变方法接受者的值
	c.radius = 10
	return area
}

//相当于引用传递
/*
	* 指针是一个变量，指针变量的值是一个内存地址
	* 因为指针本身是一个变量，所以指针变量也可以被另外一个指针指向，因此则存在指针的指针
 */
func (c *Circle) getArea2() float64 {
	//c.radius 即为 Circle 类型对象中的属性
	var area = 3.14 * c.radius * c.radius
	// 改变接受的内容
	c.radius = 10
	return area
}

func testRef(c Circle) {
	c.radius = 11
}

func testRef2(c *Circle) {
	c.radius = 12
}
