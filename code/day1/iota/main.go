package main

import "fmt"

const (
	n1 = iota
	n2 = 100
	n3 = iota
	n4
)

const (
	d1, d2 = iota + 1, iota + 2 //d1 = 1,d2 = 2 iota在const出现时置为0，每进行一行常量声明+1，d1,d2在同一行所以 iota=0
	d3, d4 = iota + 1, iota + 2 //iota = 1
)

func main() {

	fmt.Println(d1)
	fmt.Println(d2)
	fmt.Println(d3)
	fmt.Println(d4)
}
