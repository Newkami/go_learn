package main

import "fmt"

type Dog struct {
	name string
}

func newDog(name string) Dog {
	return Dog{
		name: name,
	}
}

func (d Dog) wang() { //d为形式参数
	fmt.Printf("%s汪汪汪", d.name)
}

func main() {
	dog1 := newDog("Lika")
	dog1.wang()
}

//在函数前加（）指定调用的类型，该函数就可成为方法，
//接收者表示的是调用该方法的具体类型变量
