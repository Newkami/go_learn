## Day5结构体

#### 类型别名和自定义类型

`type`关键字来定义自定义类型 ,type后面跟的是类型

```go
type MyInt int
```

类型别名

```go
type TypeAlias = Type
```

```go
//类型定义
type NewInt int

//类型别名
type MyInt = int

func main() {
	var a NewInt
	var b MyInt
	
	fmt.Printf("type of a:%T\n", a) //type of a:main.NewInt
	fmt.Printf("type of b:%T\n", b) //type of b:int
```

结果显示a的类型是`main.NewInt`，表示main包下定义的`NewInt`类型。b的类型是`int`。`MyInt`类型只会在代码中存在，编译完成时并不会有`MyInt`类型。

### 结构体

Go语言提供了一种自定义数据类型，可以封装多个基本数据类型，这种数据类型叫**结构体**，英文名称`struct`。 也就是我们可以通过`struct`来定义自己的类型了。

Go语言中通过`struct`来实现面向对象。

```go
type 类型名 struct {
    字段名 字段类型
    字段名 字段类型
    …
}
```

- 类型名：标识自定义结构体的名称，在同一个包内不能重复。
- 字段名：表示结构体字段名。结构体中的字段名必须唯一。
- 字段类型：表示结构体字段的具体类型。

比如定义一个人的结构体

```go
type Person struct{
	name,city string
	age int8
}
```

结构体只有被实例化时，才会真正的分配内存，实例化后才能使用结构体中的字段。

结构体本身也是一种类型，我们可以像声明内置类型一样使用`var`关键字声明结构体类型。

```go
var p1 Person
p1.name = "Newkami"
p1.city = "Nanjing"
p1.age = 20
```

**匿名结构体**

```go
var user struct{Name string; Age int} //同一行多条语句需要加分号
user.Name = "Newkami"
user.Age = 18
```

结构体是值类型，还可以通过`new`关键字对结构体进行**实例化**，得到的是结构体的地址

```go
var p1 = new(person)
```

Go语言对指针类型变量进行操作的话，可以不用加*，Go会自动识别该地址下的变量

##### 取结构体的地址实例化

使用`&`对结构体进行取地址操作相当于对该结构体类型进行了一次`new`实例化操作。

```go
p3 := &person{}
fmt.Printf("%T\n", p3)     //*main.person
fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"", city:"", age:0}
```

##### 结构体初始化

没有初始化的结构体，其成员变量都是对应其类型的零值。初始化方法：

##### 使用键值对初始化

```go
p5 := person{
	name: "Newkami",
	city: "南京",
	age:  18,
}
//也可以对结构体指针进行键值对初始化 p5 := &person{content}
```

##### 使用值的列表初始化

```go
p6 := person{
	"Newkami",
	"南京",
	18,
}
//必须要按照结构体定义的顺序初始化
/*必须初始化结构体的所有字段。
初始值的填充顺序必须与字段在结构体中的声明顺序一致。
该方式不能和键值初始化方式混用*/
```

结构体占用一块连续的内存

#### 构造函数：返回一个结构体变量的函数

Go语言的结构体没有构造函数，我们可以自己实现。 例如，下方的代码就实现了一个`person`的构造函数。 因为`struct`是值类型，如果结构体比较复杂的话，值拷贝性能开销会比较大，所以该构造函数返回的是结构体指针类型。构造函数一般以`new`开头

```go
func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}
```

#### 方法和接收者

Go语言中的`方法（Method）`是一种作用于特定类型变量的函数。这种特定类型变量叫做`接收者（Receiver）`。接收者的概念就类似于其他语言中的`this`或者 `self`

```go
type Dog struct
{
	name string
}

func newDog(name string) Dog{
    return Dog{
        name:name,
    }
}

func (d Dog)wang(){  //d为形式参数
    fmt.Printf("%s汪汪汪",d.name)
} 

func main(){
    dog1 := newDog("Lika")
    dog1.wang()
}
//在函数前加（）指定调用的类型，该函数就可成为方法，
//接收者表示的是调用该方法的具体类型变量
```

标识符：变量名 函数名 类型名 方法名

Go语言中如果标识符首字母是大写的，就表示对外部可见（暴露的，公有的）

#### 指针类型和值类型的接收者

```go
type person struct
{
	name string
	age int8
}

func newPerson(name string,age int) person{
    return person{
        name:name,
        age:age.
    }
}
func (p person)setAge(newAge int){
    p.age = newAge
}
func (p *person)setAge2(newAge int){
    p.age = newAge
}
func main(){
    p1 := newPerson("Newkami",18)
    fmt.Println(p1.age) //18
    p1.setAge(20) //因为p为值类型的接收者，所以修改的是它的拷贝
    fmt.Println(p1.age) //18
    p1.setAge2(20)//此处setAge2方法的接收者为指针类型的接收者，修改的是地址指向的变量
    fmt.Println(p1.age) //20
}
```

##### 什么时候应该使用指针类型接收者

1. 需要修改接收者中的值
2. 接收者是拷贝代价比较大的大对象
3. 保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。

#### 自定义类型加方法

在Go语言中，接收者的类型可以是任何类型，不仅仅是结构体，任何类型都可以拥有方法。 举个例子，我们基于内置的`int`类型使用type关键字可以定义新的自定义类型，然后为我们的自定义类型添加方法。

```go
//MyInt 将int定义为自定义MyInt类型
type MyInt int

//SayHello 为MyInt添加一个SayHello的方法
func (m MyInt) SayHello() {
	fmt.Println("Hello, 我是一个int。")
}
func main() {
	var m1 MyInt
	m1.SayHello() //Hello, 我是一个int。
	m1 = 100
	fmt.Printf("%#v  %T\n", m1, m1) //100  main.MyInt
}
```

不能给别的包里面的类型添加方法，只能给当前包的类型添加方法

#### 匿名字段

结构体允许在声名时只有类型没有字段名，被称为匿名字段

```go
//Person 结构体Person类型
type Person struct {
	string
	int
}

func main() {
	p1 := Person{
		"Newkami",
		18,
	}
	fmt.Printf("%#v\n", p1)        //main.Person{string:"Newkami", int:18}
	fmt.Println(p1.string, p1.int) //北京 18
```

匿名字段并不代表没有字段名，而是默认会采用类型名作为字段名，结构体要求字段名称必须唯一，因此一个结构体中同种类型的匿名字段只能有一个。