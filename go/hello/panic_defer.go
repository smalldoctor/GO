package main

func test_defer(num int32) func(res int32) (final int32) {
	tmp := num
	/*
	defer函数是在return语句执行之前执行的,即执行defer之后，再执行return语句
	多个defer，按照定义的顺序倒序执行;

	defer也是一个普通的执行语句，类似于 tmp := num；
	因此defer后面跟的任何变量都是立即执行的，参考值传递还是地址传递；

	self struct
	defer self.method()
	self是立即赋值，且是值赋值；后面对self的修改不会影响到这里的self；

	self *struct
	defer self.method()
	self是立即赋值，且是地址赋值；后面对self的修改会影响到这里的self；
	*/
	defer func() {
		tmp = tmp + 3
	}()
	return func(res int32) (final int32) {
		res = tmp - 3
		return res
	}
}

func main() {
	res := test_defer(10)
	print("res: ", res(10))
}
