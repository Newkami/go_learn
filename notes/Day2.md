## Day2 基本数据类型知识

##### 数字字面量语法

Go1.13版本之后引入了数字字面量语法，便于开发者以二进制、八进制或十六进制浮点数的格式定义数字

`v := 0b00101101`， 代表二进制的 101101，相当于十进制的 45。 

`v := 0o377`，代表八进制的 377，相当于十进制的 255。

`v := 0x1p-2`，代表十六进制的 1 除以 2²，也就是 0.25。

`uintptr` 无符号整型 可用于存放一个指针

**注意事项** 在涉及到二进制传输、读写文件的结构描述时，为了保持文件的结构不会受到不同编译目标平台字节长度的影响，不要使用`int`和 `uint`。

##### 浮点数

支持`float32`和`float64` `float32`的浮点数的最大范围约为`3.4e38` 

`float64` 的浮点数的最大范围约为 `1.8e308`

最大值分别可用`math.MaxFloat32`和`math.MaxFloat64`定义

##### 复数

复数有实部和虚部，`complex64`的实部和虚部为32位，`complex128`的实部和虚部为64位

##### 布尔值

1. 布尔类型变量的默认值为`false`。
2. Go 语言中**不允许**将整型强制转换为布尔型.
3. 布尔型**无法参与数值运算**，也无法与其他类型进行转换。

##### 字符串

go中字符串是双引号包裹

单引号包裹的为字符

一个字符‘A’占1个字节 一个utf8编码汉字一般占3个字节

| 转义符 |                含义                |
| :----: | :--------------------------------: |
|  `\r`  |         回车符（返回行首）         |
|  `\n`  | 换行符（直接跳到下一行的同列位置） |
|  `\t`  |               制表符               |
|  `\'`  |               单引号               |
|  `\"`  |               双引号               |
|  `\\`  |               反斜杠               |

##### 多行字符串

使用反引号字符：

```go
s1 := `第一行
第二行
第三行
`
fmt.Println(s1)
```

字符串常用操作

len(str) 求长度

+或fmt.Sprintf 拼接字符串

```go
  name := "Newkami"
  age := "18 years old"
  s := name + age

  fmt.Println(s)
  fmt.Printf("%s%s", name, age)
  s1 := fmt.Sprintf("%s%s", name, age) //使用该方法可返回一拼接后的字符串
  fmt.Println(s1)
```

strings.Split 分割

```go
age := "18,years,old"
ret := strings.Split(age, ",") //[18 years old]
```

strings.contains 判断是否包含

```go
s := "Newkami"
fmt.Println(strings.Contains(s, "New")) //true
```

strings.HasPrefix,strings.HasSuffix 前后缀判断

```go
s := "Newkami"
fmt.Println(strings.HasPrefix(s, "New")) //true
fmt.Println(strings.HasSuffix(s, "New")) //false
```

strings.Index(),strings.LastIndex() 子串出现位置

```go
s := "Newka_N_mi"
fmt.Println(strings.Index(s, "N")) //0
fmt.Println(strings.LastIndex(s, "N")) //6
```

strings.Join(a[]string, sep string) join操作

```go
 age := "18,years,old"
 ret := strings.Split(age, ",")
 fmt.Println(strings.Join(ret, "+")) //18+years+old
```

##### 字符

1. `uint8`类型，或者叫 byte 型，代表了`ASCII码`的一个字符。
2. `rune`类型，代表一个 `UTF-8字符`。

当需要处理中文、日文或者其他复合字符时，则需要用到`rune`类型。`rune`类型实际是一个`int32`

##### 修改字符串

要修改字符串，需要先将其转换成`[]rune`或`[]byte`，完成后再转换为`string`。无论哪种转换，都会重新分配内存，并复制字节数组。

##### 类型转换

Go语言中只有强制类型转换，没有隐式类型转换。该语法只能在两个类型之间支持相互转换的时候使用。

基本语法 `T(表达式)`

## 流程控制

##### if条件判断特殊写法

可以在 if 表达式之前添加一个执行语句，再根据变量值进行判断

```go
func ifDemo2() {
	if score := 65; score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}
```

score变量在if结构中使用完就消失了，并不会保存，也不能在其他地方使用

##### for range

Go语言中可以使用`for range`遍历数组、切片、字符串、map 及通道（channel）。 通过`for range`遍历的返回值有以下规律：

1. 数组、切片、字符串返回索引和值。
2. map返回键和值。
3. 通道（channel）只返回通道内的值。

##### **switch**

`fallthrough`语法可以执行满足条件的case的下一个case，是为了兼容C语言中的case设计的。

```go
func switchDemo5() {
	s := "a"
	switch {
	case s == "a":
		fmt.Println("a")
		fallthrough
	case s == "b":
		fmt.Println("b")
	case s == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}
}
```

输出为 a \n b

##### goto

`goto`语句通过标签进行代码间的无条件跳转。`goto`语句可以在快速跳出循环、避免重复退出上有一定的帮助。

## 数组

数组长度是数组类型的一部分

数组不初始化，默认值为零值（false，0，“”）

```go
var a1 [3]bool
var a2 [4]bool
a1 = [3]bool{false,true,true}
//根据初始值自动推断数组长度
a0 := [...]int{1,2,3,4,5,6,7}
a3 := [5]int{1,2} //[1,2,0,0,0]
a3 := [5]int{0:1,4:2} //[1,0,0,0,2] 根据索引初始化
```

##### 数组遍历

```go
//根据索引遍历

for i:=0;i<len(a0);i++{
	fmt.Println(a0[i])
}
//for range
for i,v:= range a0
	fmt.Printf("%d",v)
```

##### 数组是值类型

数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值。