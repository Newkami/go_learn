package main

import "fmt"

func main() {
	/*
		var s []int //定义一个存放int类型元素的切片
		var str []string

		s = []int{1, 2, 3}
		str = []string{"yi", "er", "san"}
		fmt.Println(s, str)
		fmt.Println(len(str)) //3
		fmt.Println(cap(str)) //3
	*/
	a1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	s3 := a1[3:7] //左闭右开 [4 5 6 7]
	fmt.Println(s3)
	s4 := s3[1:3]
	fmt.Println(s4)      // [5 6]
	fmt.Println(len(s4)) //2
	fmt.Println(cap(s4)) //4  s4 是基于 s3切片的 但底层数组还是由a1决定

}
