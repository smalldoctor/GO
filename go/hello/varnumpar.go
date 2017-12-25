package main

import "fmt"

func main() {
	x := min(1, 3, 2, 0)
	fmt.Printf("The minimum is: %d\n", x)
	arr := []int{7, 9, 3, 5, 1}
	// 作为可变长参数传递
	x = min(arr...)
	fmt.Printf("The minimum in the array arr is: %d", x)

	cstr := make(chan string)
	fmt.Println(cstr)
}

/*
...int是一个可变长参数；
min是一个可以接受可变个数参数的函数
*/
func min(a ...int) int {
	if len(a) == 0 {
		return 0
	}
	min := a[0]
	for _, v := range a {
		if v < min {
			min = v
		}
	}

	return min
}

/*

但是如果变长参数的类型并不是都相同的呢？使用 5 个参数来进行传递并不是很明智的选择，有 2 种方案可以解决这个问题：

使用结构：

定义一个结构类型，假设它叫 Options，用以存储所有可能的参数：

  type Options struct {
      par1 type1,
      par2 type2,
      ...
  }
函数 F1 可以使用正常的参数 a 和 b，以及一个没有任何初始化的 Options 结构： F1(a, b, Options {})。
如果需要对选项进行初始化，则可以使用 F1(a, b, Options {par1:val1, par2:val2})。

使用空接口：

如果一个变长参数的类型没有被指定，则可以使用默认的空接口 interface{}，这样就可以接受任何类型的参数。
该方案不仅可以用于长度未知的参数，还可以用于任何不确定类型的参数。使用一个 for-range 循环以及 switch 结构对每个参数的类型进行判断：

  func typecheck(..,..,values … interface{}) {
      for _, value := range values {
          switch v := value.(type) {
              case int: …
              case float: …
              case string: …
              case bool: …
              default: …
          }
      }
  }
*/
