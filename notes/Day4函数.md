## Day4函数

函数定义

```go
func 函数名(参数)(返回值){
    函数体
}
```

go语言是强类型语言，一个变量必须要指定类型，所以参数也要定义类型

```go
func sum(x int, y int)(ret int)
{
    return x+y
}
```

函数是一段代码的封装，把一段逻辑抽象出来封装到一个函数中，使代码结构更清晰更整洁

函数可以没有参数，也可以没有返回值

返回值可以命名也可以不命名

```go
func f1()int
{
    ret:=3
    return 3
}
```

命名的返回值就相当于在函数中声明一个变量,显式命名之后，`return`可以省略

```go
func sum(x int, y int)(ret int)
{
    ret = x + y
}
```

**多个返回值**

```go
func f2()(int string)
{
    return 1,"Newkami"
}
```

参数的**类型简写**

```go
func f3(x, y, z int,m, n string)(ret int){}
```

**可变长参数**

可变长参数必须放在函数参数的最后

```go
func f3(s string,y...int)
{
    fmt.Println(x)
    fmt.Println(y)
}
f3("Newkame",1,2,3,4,5) //"Newkami" [1 2 3 4 5]返回的是一个切片
```

Go中没有默认参数这个概念

### defer语句

Go语言中的`defer`语句会将其后面跟随的语句进行延迟处理。在`defer`归属的函数即将返回时，将延迟处理的语句按`defer`定义的逆序进行执行，也就是说，先被`defer`的语句最后被执行，最后被`defer`的语句，最先被执行。

```go
func main() {
	fmt.Println("start")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end")
}  

/*输出结果 
start
end
3
2
1*/
```

##### defer执行时机

在Go语言的函数中`return`语句在底层并不是原子操作，它分为给返回值赋值和RET指令两步。而`defer`语句执行的时机就在返回值赋值操作后，RET指令执行前。

```go
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))  //先算出函数内部的函数值
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
}
/* 
1.defer calc("AA", x, calc("A", x, y))
2.计算calc("A", x, y) x = 1 y = 2  //A 1 2 3
3.def clac("AA", 1 , 3)
4.x = 10
5.defer calc("BB", x, calc("B", x, y))
6.计算calc("B", x, y) x = 10,y = 2 //B 10 2 12
7.def calc("BB", 10, 12) x = 10
8.y = 20
9.计算 calc("BB", 10, 12) // BB 10 12 22
10.计算 clac("AA", 1 , 3) // AA 1 3 4
```

defer注册要延迟执行的函数时该函数所有的参数都需要确定其值，也就是说defer前，函数内部的参数值都要确定好，不会因后面变量值发生改变而改变。

### 函数进阶

#### **变量作用域**

函数中查找变量的顺序：

1.先在函数内部查找

2.在函数外部查找，一直找到全局变量

全局变量是定义在函数外部的变量，它在程序整个运行周期内都有效。 在函数中可以访问到全局变量。

局部变量又分为两种： 函数内定义的变量无法在该函数外使用；语句块定义的变量，通常我们会在if条件判断、for循环、switch语句上使用这种定义变量的方式。

#### 函数类型

函数也是一种类型，函数也可作为参数的类型

```go
func f1() {}  //%T 的值为func ()
func f2()int {}  //%T 的值为func () int
func f3(x func() int) {
    ret := x()
    fmt.Println(ret)
}
```

函数也可以作为返回值

```go
func f4(x func() int)func(int, int) int {}
//f4的返回值是一个func(int,int) int类型
```

#### 匿名函数

没有名字的函数，可以定义在函数内部，匿名函数可以用变量接收，在函数内部无法定义带名的函数

如果只是调用一次的函数，还可以简写成立即执行函数

```go
func main(){
	func(x,y int){
        fmt.Println(x+y)
    }(100,200) //扩号代表定义完成后立即执行函数，不存储该函数变量
}
```

#### 闭包

闭包指的是一个函数和与其相关的引用环境组合而成的实体。简单来说，`闭包=函数+引用环境`

```go
func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func adder2(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

func main() {
    var f = adder()  
    // f是一个func(int) int的变量 f这个变量可以调用adder()中返回值
    // 也就是 adder return的匿名函数
    // f也就是是一个有一个int型参数 返回值为int型的函数，被调用时执行匿名函数的语句
    // 闭包等于函数加上外部变量的引用
	fmt.Println(f(10)) //10
	fmt.Println(f(20)) //30
	fmt.Println(f(30)) //60

	f2 := adder2(0)
	fmt.Println(f1(40)) //40
	fmt.Println(f1(50)) //90
}
```

变量`f`是一个函数并且它引用了其外部作用域中的`x`变量，此时`f`就是一个闭包。 在`f`的生命周期内，变量`x`也一直有效。

```go
//例子：要求在f1中执行f2
func f1(f func())
{
	fmt.Println("this is f1")
    f()
}

func f2(x,y int){
	fmt.Println("this is f2")
	fmt.Println(x+y)
}
//定义一个f3返回值为f1参数类型，参数为f2的函数类型
func f3(f func(int int),x,y int)func (){
    tmp:=func()
    {
        f(x,y) //实际上执行的为f2(x,y)
    }
    return tmp
}
//f3就是一个闭包，它引用了外部环境的x,y变量来执行f2(x,y)
```

### 内置函数

|    内置函数    |                             介绍                             |
| :------------: | :----------------------------------------------------------: |
|     close      |                     主要用来关闭channel                      |
|      len       |      用来求长度，比如string、array、slice、map、channel      |
|      new       | 用来分配内存，主要用来分配值类型，比如int、struct。返回的是指针 |
|      make      |   用来分配内存，主要用来分配引用类型，比如chan、map、slice   |
|     append     |                 用来追加元素到数组、slice中                  |
| panic和recover |                        用来做错误处理                        |

1. `recover()`必须搭配`defer`使用。
2. `defer`一定要在可能引发`panic`的语句之前定义。