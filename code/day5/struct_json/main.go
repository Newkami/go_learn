package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name"` //字段名必须要大写才能被其他包识别到
	Age  int
}

func main() {
	p1 := person{
		Name: "Newkami",
		Age:  18,
	}
	b, err := json.Marshal(p1) //序列化
	if err != nil {
		fmt.Println("marshal failed:%v", err)
		return
	}
	fmt.Printf("%#v\n", string(b))
	str := `{"name":"Newkami","age":18}`
	var output person
	json.Unmarshal([]byte(str), &output) //传指针为了能在Unmarshal内部修改output的值
	fmt.Printf("%#v", output)
}
