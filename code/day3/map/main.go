package main

import "fmt"

func main() {
	var sliceMap = make(map[string][]int, 3)
	sliceMap["ages"] = []int{18, 28, 40}
	fmt.Println(sliceMap)
}
