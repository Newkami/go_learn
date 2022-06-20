DAY3

### 切片（slice）

切片（Slice）是一个拥有相同类型元素的可变长度的序列。它是基于数组类型做的一层封装。它非常灵活，支持自动扩容。

切片是一个引用类型，它的内部结构包含`地址`、`长度`和`容量`。切片一般用于快速地操作一块数据集合。

len()求长度, cap()求容量

```go
a1 := [...]int{1, 2, 3, 4, 5, 6}
s3 := a1[0:4] //左闭右开 [1 2 3 4]
s5 := a1[:4] //=>[0:4] [1 2 3 4]
//len(s5)=4 cap(s5) = 6 切片指向的是底层的数组
s6 := a1[3:] //=>[3:len(a1)] [4 5 6]
//len(s6) = cap(s6) = 3 切片的容量是底层数组从切片的第一个元素到最后的元素数量
```

切片没有自己的值 切片指向的底层数组值发生改变，切片的值也会发生改变

```go
a1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
s3 := a1[3:7] //左闭右开 [4 5 6 7]
fmt.Println(s3)
s4 := s3[1:3]
fmt.Println(s4)    // [5 6]
fmt.Println(len(s4)) //2
fmt.Println(cap(s4)) //4  s4 是基于 s3切片的 但底层数组还是由a1决定
```

##### make()构造切片

```go
s1 := make([]int,5,10)
fmt.Print("s1=%v len(s1)=%d cap(s1)=%d",s1,len(s1),cap(s1)) 
//s1=[0 0 0 0 0] len(s1)=5 cap(s1) = 10
```

切片的本质就是一段连续的内存。属于引用类型，实际数据存在底层数组中

**切片之间不能比较**，不能使用`==`判断切片是否相等，只能和`nil`进行比较，一个`nil`值的切片没有底层数组，长度和容量都为0

##### 切片的遍历 

同样支持索引遍历和range遍历

##### 为切片追加元素

`slice=append(slice, content)`

调用append的函数必须用原来的切片接收返回值，append函数会将放不下的底层数组换一个，也就是底层数组的地址发生了变化，所以一定要用变量接收返回值

**要检查切片是否为空**，请始终使用`len(s) == 0`来判断，而不应该使用`s == nil`来判断。

**copy()函数拷贝切片**

```go
copy(destSlice, srcSlice []T)
```

copy得到的切片和源切片不共享底层数组

切片没有删除元素的方法

要从切片a中删除索引为`index`的元素，操作方法是`a = append(a[:index], a[index+1:]...)`

```go
func main() {
	// 从切片中删除元素
	a := []int{30, 31, 32, 33, 34, 35, 36, 37}
	// 要删除索引为2的元素
	a = append(a[:2], a[3:]...)
	fmt.Println(a) //[30 31 33 34 35 36 37]
}
```

### 指针

Go语言中不存在指针操作

1.`&`：取地址

```go
n:=18
fmt.Println(&n)
```

2.`*`：根据地址取值

```go
p:= &n
m = *p
fmt.Println(m) //18
```

指针定义时值为`nil` ，若需要定义指针使用`new`函数申请一个内存地址

```go
var a = new(int)
```

`make`也用于内存分配，他只用于slice，map，chan的内存创建，返回的类型就是这三个类型本身，而不是指针类型，区别于`new`，很少用，一般用于给基本数据类型申请内存。

### map

map是一种无序的基于`key-value`的数据结构，Go语言中的map是引用类型，必须**初始化**才能使用。

```go
map[KeyType]ValueType
```

判断键是否存在

```go
value, ok := map[key]
```

map的**遍历**使用for range遍历

使用`delete()`内建函数从map中删除一组键值对，`delete()`函数的格式如下：

```go
delete(map, key)
```

### 元素为map类型的切片

```go
var mapSlice = make([]map[string]string, 3)
```

切片中的map并没有被初始化，对map初始化后才能操作

```go
mapSlice[0] = make(map[string]string, 10)
mapSlice[0]["name"] = "Newkami"
mapSlice[0]["password"] = "123456"
```

### 值为切片类型的map

```go
var sliceMap = make(map[string][]int, 3)
sliceMap["ages"] = []int{18,28,40}
fmt.Println(sliceMap)
```

##### 按照指定顺序遍历map

```go
func main() {
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子

	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		scoreMap[key] = value
	}
	//取出map中的所有key存入切片keys
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}
```