## Day6

### 接口

接口是一种类型！！！一种抽象的类型。

接口类型更注重“我能做什么”的问题。接口类型就像是一种约定——概括了一种类型应该具备哪些方法，在Go语言中提倡使用面向接口的编程方式实现解耦。

```go
type speaker interface{
    speak()  //只要实现了speak()方法的变量都是speaker类型
}
```

#### 接口的定义

每个接口类型由任意个方法签名组成，接口的定义格式如下：

```go
type 接口类型名 interface{
    方法名1( 参数列表1 ) 返回值列表1
    方法名2( 参数列表2 ) 返回值列表2
    …
}
```

#### 接口的实现

一个变量如果实现了接口中规定的所有方法，那么这个变量就实现了这个接口，可以称为这个接口类型的变量

实现的方法的参数，返回值都必须相同才算真正实现

接口的结构分为类型和值两部分，可以动态的接收类型和值，接口是引用类型

#### 值接收者和指针接收者

实现接口时使用值接收者和使用指针接收者区别

```go
type animal interface{
    move()
}

type cat struct{
    name string
    feet int8
}

func(c cat)move(){  //使用值接收者实现接口
    fmt.Println("猫走路")
}


func main(){
    var a1 animial
    c1 := cat{"tom",4}
    c2 := &cat{"Kami",4}
    a1 = c1
    fmt.Println(a1)
    a1 = c2
    fmt.Println(a1)  //&{Kami,4}
}
//————————————————————————————————————
//和上面区分下
func(c *cat)move(){  //使用指针接收者实现接口
    fmt.Println("猫走路")
}

func main(){
    var a1 animial
    c1 := cat{"tom",4}
    c2 := &cat{"Kami",4}
    //error a1 = c1   //此时不能用值类型赋值
    fmt.Println(a1)
    a1 = c2
    fmt.Println(a1)  //&{Kami,4}
}

```

使用值接收者，结构体类型和结构体指针类型的变量都能存

使用指针接收者，只有**结构体指针类型**的变量能存

#### 接口和类型

多个结构体类型可以实现同一个接口

一个类结构体型也可以实现多个接口

接口与接口之间可以通过互相嵌套形成新的接口类型，例如Go标准库`io`源码中就有很多接口之间互相组合的示例。

```go
// src/io/io.go

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

// ReadWriter 是组合Reader接口和Writer接口形成的新接口类型
type ReadWriter interface {
	Reader
	Writer
}

// ReadCloser 是组合Reader接口和Closer接口形成的新接口类型
type ReadCloser interface {
	Reader
	Closer
}

// WriteCloser 是组合Writer接口和Closer接口形成的新接口类型
type WriteCloser interface {
	Writer
	Closer
}
```

对于这种由多个接口类型组合形成的新接口类型，同样只需要实现新接口类型中规定的所有方法就算实现了该接口类型。

### 空接口

空接口中没有定义任何函数，通常按以下方式定义

```go
interface {}
```

说明任何类型都实现了这个接口，空接口可用于函数中接收任意类型的参数

```go
//interface是关键字
//interface{}是空接口类型

func main(){
	var m1 map[string]interface{}
    m1["name"] = "Newkami"
    m1["age"] = 20
    m1["married"] = true
} //实现了map值存放类型不确定的问题
```

#### 类型断言

接口值可能赋值为任意类型的值，那我们如何从接口值获取其存储的具体数据呢

而想要从接口值中获取到对应的实际值需要使用**类型断言**，其语法格式如下。

```go
x.(T)
```

其中：

- x：表示接口类型的变量
- T：表示断言`x`可能是的类型。

该语法返回两个参数，第一个参数是`x`转化为`T`类型后的变量，第二个值是一个布尔值，若为`true`则表示断言成功，为`false`则表示断言失败。

```go
var n Mover = &Dog{Name: "旺财"}
v, ok := n.(*Dog)
if ok {
	fmt.Println("类型断言成功")
	v.Name = "富贵" // 变量v是*Dog类型
} else {
	fmt.Println("类型断言失败")
}
```

```go
// justifyType 对传入的空接口类型变量x进行类型断言
func justifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}
```

只有当有两个或两个以上的具体类型必须以相同的方式进行处理时才需要定义接口。切记不要为了使用接口类型而增加不必要的抽象，导致不必要的运行时损耗。