package main

import (
	"fmt"
)

func changeString() {
	s1 := "big"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1))

	s2 := "白萝卜"
	runeS2 := []rune(s2)
	runeS2[0] = '红'
	fmt.Println(string(runeS2))
}
func main() {
	/*
		name := "Newkami"
		age := "18,years,old"
		s := name + age
		fmt.Println(s)
		fmt.Printf("%s%s", name, age)
		s1 := fmt.Sprintf("%s%s", name, age) //使用该方法可返回一拼接后的字符串
		fmt.Println(s1)
		ret := strings.Split(age, ",") //返回的是类似于列表的类型
		fmt.Println(ret)
	*/
	// s := "Newkami"
	// fmt.Println(strings.Contains(s, "New"))
	// s := "Newka_N_mi"
	// fmt.Println(strings.Index(s, "N"))
	// fmt.Println(strings.LastIndex(s, "N"))
	// age := "18,years,old"
	// ret := strings.Split(age, ",")
	// fmt.Println(strings.Join(ret, "+"))
	changeString()
}
