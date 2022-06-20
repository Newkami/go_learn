package main

import "fmt"

//经典面试题
func f1() int {
	x := 5
	defer func() {
		x++ //此处修改的是x，不是返回值
	}()
	return x //先把x的值赋值给RET，再返回RET
}

func f2() (x int) { //使用的命名返回值，返回值就是x
	defer func() {
		x++
	}()
	return 5 //把5赋值给返回值x
}

func f3() (y int) { //使用的命名返回值，返回值是y
	x := 5
	defer func() {
		x++ //修改的是x
	}()
	return x // y = x 再 return y

}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x) //修改的是函数中的副本x
	return 5 //RET = x =5
}
func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}
