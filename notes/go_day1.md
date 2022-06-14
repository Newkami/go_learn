Let’s Go!

## go环境搭建

https://golang.google.cn/dl/  国内go开发包的镜像网站 下载并安装

打开命令行

```
go version #检查安装是否完成
```

##### 配置GOPATH

GOPATH是一个环境变量 相当于配置一个工作目录

任意找一个目录添加为系统环境变量 GOPATH

建立src ，pkg，bin文件夹

##### GO项目结构

个人开发者 src里放各个项目的文件夹

比如 src->github.com->code内容

##### 下载vscode

安装chinese中文插件 和 go插件



##### 编写第一个go程序

```go
package main
import "fmt"

func main() {
  fmt.Println("Hello world")
}
```

##### 编译

1.在项目文件夹下使用 `go build`

得到一个exe可执行文件 就可以直接运行了

2.在其他路径下go build 需要在后面加上项目的路径 从 GOPATH/src 往后写

3.`go build -o name.exe 自定义文件名`

类似有 `go run` 和 `go install` 命令

**go支持跨平台编译**（windows平台）

在windows中编译一个linux平台运行的文件

```com
SET CGO_ENABLED=0  // 禁用CGO
SET GOOS=linux  // 目标平台是linux
SET GOARCH=amd64  // 目标处理器架构是amd64
```

Linux平台下编译Windows平台64位可执行程序

```bash
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
```



## GO文件基本结构

1.`package` 声明该文件属于哪个包 `package main` 表明可编译为exe文件，其中必须有`func main()`函数作为程序入口才能编译 main函数没有参数和返回值

2.import为包导入关键字 使用“ ”包括

3.函数外只能放置**变量 常量 函数 类型**的声明，具体的语句不能放置在函数外

例如 `fmt.Println()`不能放置在函数外

## 变量和常量

#### 变量的来历

程序运行过程中的数据都是保存在内存中，我们想要在代码中操作某个数据时就需要去内存上找到这个变量，但是如果我们直接在代码中通过内存地址去操作变量的话，代码的可读性会非常差而且还容易出错，所以我们就利用变量将这个数据的内存地址保存起来，以后直接通过这个变量就能找到内存上对应的数据了。

Go语言的变量声明格式为：

```go
var 变量名 变量类型
```

例如：

```go
var name string
var age int
var isOk bool
```

go语言支持**批量声明**(避免过多的var关键字使用)

```go
var (
    a string
    b int
    c bool
    d float32
)
```

**变量初始化**和其他语言大同小异，可以一次性初始化多个变量

```go
var name, age = "Newkami", 20
```

**类型推导**

有时候会将变量类型省略，编译器会自动推导变量类型

**短变量声明**

在函数内部，可以使用更简略的 `:=` 方式声明并初始化变量。

在使用多重赋值时，如果想要忽略某个值，可以使用`匿名变量（anonymous variable）`。 匿名变量用一个下划线`_`表示，例如：

```go
func foo() (int, string) {
	return 20, "Newkami"
}
func main() {
	x, _ := foo()
	_, y := foo()
	fmt.Println("x=", x)
	fmt.Println("y=", y)
}
```

匿名变量不占用命名空间，不会分配内存，所以匿名变量之间不存在重复声明。

**注意事项：**

1. 函数外的每个语句都必须以关键字开始（var、const、func等）
2. `:=`不能使用在函数外。
3. `_`多用于占位，表示忽略值。

#### 常量

常量的定义指不变的值，相对于变量只是把关键字`var`换成了`const`，常量在定义的时候必须赋值。

const同时声明多个常量时，如果省略了值则表示和上面一行的值相同。 例如：

```go
const (
    n1 = 100
    n2
    n3
)
```

上面示例中，常量`n1`、`n2`、`n3`的值都是100。

#### iota

`iota`是go语言的常量计数器，只能在常量的表达式中使用。

`iota`在const关键字出现时将被重置为0。const中每新增一行**常量声明**将使`iota`计数一次(iota可理解为const语句块中的行索引)。 使用iota能简化定义，在定义枚举时很有用。

```go
const (
		n1 = iota //0
		n2        //1
		n3        //2
		n4        //3
	)
```

定义数量级 （这里的`<<`表示左移操作，`1<<10`表示将1的二进制表示向左移10位，也就是由`1`变成了`10000000000`，也就是十进制的1024。同理`2<<2`表示将2的二进制表示向左移2位，也就是由`10`变成了`1000`，也就是十进制的8。）

```go
const (
		_  = iota
		KB = 1 << (10 * iota)
		MB = 1 << (10 * iota)
		GB = 1 << (10 * iota)
		TB = 1 << (10 * iota)
		PB = 1 << (10 * iota)
	)
```
