package main

import "fmt"

// 批量声明 变量后面为默认值
var (
	name string //""
	age  int    //0
	isOK bool   //false
)

func main() {
	name = "Newkami"
	age = 20
	isOK = true
	//Go语言变量声明必须使用，不使用就无法编译
	fmt.Printf("name:%s", name)
	fmt.Println("age:%d", age) //Println打印完指定的内容会自动加一个换行
	fmt.Print(isOK)
}
