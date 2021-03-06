
# 参考

go web编程 

https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/preface.md

go 相关电子书 

https://github.com/EbookFoundation/free-programming-books/blob/master/free-programming-books-zh.md#golang


---



go 是一门编译型语言，静态类型语言

将源代码和依赖一起打包，进行静态编译生成本地指令

go的代码是通过package 来组织的

go不支持继承

go的并发有两种 协程、channel,使用消息传递来共享内存。


---

# 名称

- 开头是字母或下划线 
- 关键字不能使用，预置函数不能使用

如果第一个字母的大小写决定是否对包外的可见，包名本身总是小写字母

一般都是驼峰式命名风格

四个主要的声明：var 、 const 、 type 、func


```
var name type = expression
name := expression  // name 的类型由 expression类型决定 ，声明和初始化局部变量

：=  表示声明
= 赋值

```

# 指针

变量是存储值的地方

指针的值是一个变量的地址，使用指针，在不知道变量名的情况下，可以直接或间接的读取或更新变量的值

一个指针指示值所保存的位置，不是所有的值都有地址，但是所有的变量都有。





```
var a int  = 1 
  p = &a  // 表示变量a的地址 ，它的类型是（*int），& 取地址操作符 ， &a 表示获取一个指向整型变量的指针
 *p   //  其实就是使用a 的地址来获取a 的值，也就是 = 1 
```
函数返回局部变量的地址是非常安全的

指针包含变量的地址，传递一个指针给函数，函数可以间接的通过指针修改变量的值

```
& 仅用于生成其操作数对应的地址，也就是用于生成指针会出现在两个内容上：
一个是类型，* Type 这样的格式代表了一个指针类型 
一个是指针，* Pointer 这样的格式用于获取指针所对应的基本值
```

# new 函数

new(T)  //创建一个T类型的变量

每一次调用new返回一个具有唯一地址的变量，地址类型为*T

new是一个预声明的函数不是关键字

```
	p := new(int) //  *int 型的p ，指向未命名的int变量
	*p = 2  // 未命名的int 变量赋值为2 
    // 相当于一个匿名变量
```



# 变量的生命周期
包级别变量的生命周期是整个程序的执行时间

局部变量有一个动态的生命周期：

每次执行声明语句时创建一个新的实体，变量一直生存到它变得不可访问，这时它占用的存储空间被回收。

函数的参数和返回值也是局部变量，他们在闭包函数被调用的时候创建。

变量的生命周期是通过它是否可达来确定的，

var 生命的变量使用的是堆空间，虽然在函数内部但是还可以在其他地方访问

这种情况称为逃逸，变量的每一次逃逸都需要一次额外的内存分配过程

赋值只有在值对于变量类型是可赋值的时才合法


# 类型声明 

type 声明定义一个新的命名类型
```
type name underlying-type 

```

类型的声明通常出现在包级别

对于每个类型T，都有一个对应的类型转换操作T(x)将值x转换为类型T

从字符串转换成字节 slice 会分配一份字符串数据副本。



# 包和文件

每一个包给它的声明提供独立的命名空间

包可以通过控制变量在包外面的可见性或导出情况来隐藏信息

包的初始化从初始化包级别的变量开始，这些变量按照声明顺序初始化

如果包由多个go文件组成，初始化按照编译器收到文件的顺序进行，
go工具会在调用编译器前将go文件排序

这个init函数不能被调用和被引用，程序启动的时候，init函数按照它们声明的顺序自动执行


声明的作用域是指用到声明时所声明名字的源代码段


# 整数

不能隐式转换，需要显示的转换

大类型向小类型转换时，不会报错，会溢出，

- 有符号整数：int8    int16   int32   int64 
- 无符号整数：uint8   uint16  uint32  uint64

声明变量默认的是int类型  fmt.Printf    %T 可以查看该变量的类型，  unsafe.Sizeof 查看变量的占用字节数

rune类型是int32类型的同义词，指明一个值是unicode码

byte是uint8的同义词，强调一个值是原始数据，而非量值

uintptr 无符号整数，大小不明确，但足以完成存放指针，仅仅用于底层编程

有符号整数以补码表示，保留最高位作为符号位

&   AND

|   OR         

^   XOR

&^  位清空AND NOT

<<  左移
**>>**  右移


#  浮点数
float32     float64


# 复数 
complex64   complex128分别由float32和float64构成

内置的complex函数根据给定的实部和虚部创建复数

内置的real函数和imag函数则分别提取复数的实部和虚部

 
```go
package main
import "fmt"
func f()  {
  x:=complex(1,2)
 	fmt.Println(x) //(1+2i)
 
 	y :=complex(2,3)
 	fmt.Println(y) //(2+3i)
 
 	fmt.Println(x * y ) //(-4+7i)
 
 	fmt.Println(real(x*y)) //-4
 
 	fmt.Println(imag(x*y)) // 7
}
  
```

浮点数或十进制整数后面紧接着写字母i，它就会变成一个虚数


# 字符串

字符串中第i个字节不一定就是第i个字符

非ASCII字符需要两个字节

下标越界，将触发宕机异常

字符串值无法改变，字符串值本身所包含的字节序列永不可变。

不可变意味着两个字符串能安全地共用同一段底层内存，使得复制任何长度字符串的开销都少

反引号可以原封不动的输出文本的内容


# slice 

表示一个拥有相同类型元素的可变长度的序列，是一种轻量级的数据结构，可以访问数组的部分或者全部元素，这个数组称为slice的底层数组。

指针：指向数组的第一个可以从slice中访问的元素

长度：指slice中元素的个数，不能超过slice的容量

容量：从slice的起始元素到底层数组的最后一个元素间元素的个数

len和cap用来返回slice的长度和容量

一个底层数组可以对应多个slice，slice可以引用数组的任何位置，彼此之间的元素还可以重叠

如果slice的引用超过了被引用对象的容量，cap（s），那么导致程序宕机

slice比较不可以直接使用==，slice的元素是非直接的，可能slice可以包含它自身，如果底层数组元素改变，同一个slice在不同的时间就会有不同的值

不能作为map的键，但是可以==nil ，nil的slice没有对应的底层数组

make([]T ,len）  // slice 只引用了前len个元素

make([]T ,len,cap)  // 创建一个无名数组，并返回了它的一个slice

slice虽然底层数组的元素是间接引用的，但是指针、长度、容量不是


# 结构体

结构体的值可以作为参数传递给函数或者作为函数的返回值，由于效率问题，使用结构体指针的方式直接传递给函数

---


# 函数

```
func name(parameter-list type ) (result-type ){
}
```
函数声明包含  名字、形参列表、返回列表

形参列表：指定了参数名和参数类型

返回列表：指定了返回值的类型

如果几个形参或者返回值的类型相同，那么类型只需要写一次

函数的类型称为函数签名。当两个函数拥有相同的形参列表和返回列表时，认为这两个函数的类型或签名是相同的。


一个函数如果有命名的返回值，可以省略return语句的操作数，称为裸返回

裸返回是将每个命名返回结果按照顺序返回的快捷方法


# 错误

error 是内置的接口类型，

go语言的异常只是针对程序bug导致的预料外的错误，而不能作为常规的错误处理方法出现在程序中。

异常会陷入带有错误消息的控制流去处理它，通常会导致预期外的结果，错误会难以理解的栈跟踪信息报告给最终用户

当一个函数调用返回一个错误时，调用者应当负责检查错误并采取合适的处理应对


- 常见的情形是将错误传递下去
- 对于不固定或者不可预测的错误，在短暂的间隔后对操作进行重试是合乎情理的，超出一定的重试次数和限定的时间后再报错退出。
- 依旧不能顺利进行下去，调用这能够输出错误然后优雅的停止程序。

调用log.Fatalf 默认会将时间和日期作为前缀添加到错误信息前。



进程错误检查之后，检测到失败的情况往往都在成功之前。

如果检测到失败导致函数返回，成功的逻辑一般不会放在else块中而是放在外层的作用域中，

通常都会在开头有一连串的检查用来返回错误，之后跟着实际的函数体一直到最后。



# 函数变量

函数变量也有类型，而且他们可以赋给变量或者传递或者从其他函数中返回。函数变量可以像其他函数一样调用。

函数类型的零值是nil，调用一个空的函数变量将导致宕机

函数变量不仅将数据进行参数化，还将函数的行为当作参数进行传递



func关键字后面没有函数的民初，他是一个表达式，它的值称作匿名函数。

里层的函数可以使用外层函数中的变量

```go
package main

import "fmt"

func main() {
	f:=squares()
	fmt.Println(f()) // 1
	fmt.Println(f()) //4 
	fmt.Println(f()) // 9 

}

func squares()  func() int {
	var x int
	return func() int {
		x ++
		return x * x
	}
}
```

里层的匿名函数能够获取和更新外层函数的局部变量，

这些隐藏的变量引用就是我们把函数归类为引用类型而且函数变量无法进行比较的原因

函数变量类似与使用闭包方法实现的变量

变量的生命周期不是由它的作用域所决定的

循环变量的作用域的规则限制。循环里创建的所有函数变量共享相同的变量，一个可访问的存储位置，而不是固定的值。




# 延迟函数调用

defer机制就是一个普通的函数或方法调用，在调用前加上关键字，无论正常还是异常，实际调用推迟到包含defer语句的函数结束后才执行。

defer语句没有限制使用次数，

经常使用于成对的操作，比如打开、关闭。

也可以用来调试一个复杂的函数，在函数的入口和出口设置调试行为。

延迟执行的匿名函数能够改变外层函数返回给调用者的结果，延迟的函数不到函数的最后一刻是不会执行的


# 宕机

一个典型的宕机发送时，正常的程序执行会终止，goroutine中的所有延迟函数会执行，然后程序会异常退出并留下一条日志消息。



recover函数在延迟函数的内部调用，而且这个包含defer语句的函数发生宕机，recover会终止当前的宕机状态并返回宕机的值。

函数不会从之前宕机的地方继续运行而是正常返回。



# 方法 

方法是某种特定类型的函数。面向对象编程就是使用方法来描述每个数据结构的属性和操作。

方法的声明和普通函数类似，只是在函数名字前面多了一个参数。这个参数把这个方法绑定到这个参数对应的类型上

```go
package main

import "fmt"

func main() {

	point1 := Point{1, 2}
	point2 := Point{11, 22}

	fmt.Println(Add(point1, point2))
	fmt.Println(point1.Add(point2)) // 直接调用 该类型的方法
}

type Point struct {
	X, Y int
}

// 普通的函数
func Add(p1, p2 Point) (int, int) {
	return p1.X + p2.X, p1.Y + p2.Y
}

// Point 类型的方法 
func (p Point) Add(p2 Point) (int, int) {
	return p.X + p2.X, p.Y + p2.Y
}

```

第一个Add 是包级别的函数，第二个类型级别的方法 Point.Add

每个类型有自己的命名空间，可以将方法绑定到任何类型上。指针类型也可以

类型拥有的所有方法名都必须是唯一的，但不同的类型可以使用相同的方法名。


# 指针接受者的方法

由于主调函数会复制每一个实参变量，如果函数需要更新一个变量，或者如果一个实参太大而我们希望避免复制整个实参，
因此我们必须使用指针来传递变量的地址。





# 接口值

一个接口类型的值其实有两个部分：一个具体类型和该类型的一个值


# 协程

本质上是一种用户状态线程，不需要操作系统来进行抢占式调度

函数调用前加上一个go关键字，这次调用就会在一个新的goroutine中并发执行，函数返回时，这个 goroutine 也自动结束

如果这个函数有返回值，那么这个返回值将会被丢弃

main 函数返回时，程序退出，并不等待其他goroutine



# 消息通信机制channel

不要通过共享内存来通信， 而应该通过通信来共享内存

可以使用channel在多个goroutine之间传递消息

一个channel只能传递一种类型的值，这个类型需要在声明channel时指定


# select

unix中，通过调用select（）函数来监控一些列的文件句柄，一旦其中一个文件句柄发生了IO动作，该select（）调用就会被返回

go语言直接在语言级别支持select关键字，用于处理异步IO问题

```
select {
	case <-chanint:
		fmt.Println(<-chanint)
	case chanint <- 2:
		fmt.Println("写入数据")
	default:
		fmt.Println("default")
	}
	
```



# 缓冲机制 

就是在make时将缓冲区的大小作为第二个参数传入即可，缓冲区写满之前都不会阻塞，可以使用range来从缓冲区读取数据

当写数据，发现channel 已满，或者读数据发现chan为空，没有正确处理都会发生goroutine锁死

select只要其中一个case完成，程序就会继续执行。


# 同步锁

sync.Mutex ：获取锁后，其他线程只能等待

sync.RWMutex：单写多读模型，读锁占用的情况下，会阻止写，不会阻止读


# 全局唯一性操作

Once类型来抱枕全局的唯一性操作

once的Do方法可以保证全局范围内只调用指定的函数一次，而且所有其他goroutine在调用到此语句时，将会先被阻塞，直至全局唯一的once.Do调用结束后才继续




# 网络编程

```
Dial(net,addr string )( Conn,error)
```

# RPC

```
func (t *T )MethodName(T t1,replyType *t2) error
```
一个PRC服务端可以注册多个不同类型的对象，但不允许注册同一类型的多个对象

一个对象中需要满足的条件：

- 必须是在对象外部可以公开调用的方法，首字母大写
- 必须有两个参数，且参数的类型都必须是包外部可以访问的类型或者是go内建支持的类型
- 第二个参数必须是一个指针
- 方法必须返回一个error类型的值

第一参数表示返回给rpc客户端传入的参数

第二个参数表示返回给prc客户端的结果


# 协程 

**并非特指go**
轻量级线程

- 能够在单一的系统线程中模拟多个任务的并发执行
- 在一个特定的时间，只有一个任务在运行，并非真正的并行
- 被动的任务调度方式，即任务没有主动抢占时间片的说法，当一个任务正在执行时，外部没有办法中止它。要进行任务切换，只能通过由该任务自身调用yield()来主动出让cpu使用权。
- 每个协程都有自己的堆栈和局部变量
  
每个协程包含三种运行状态：挂起、运行、停止。

挂起则表示该协程尚未执行完成，但出让了时间片，以后有机会时由调度器继续执行


go中的协程可能参考了libtask




---

bee执行流程

```
1、main文件监听8080 
2、用户新连接 
3、初始化context 判断类型，WebSocket？
4、执行过滤器 判断是否设置过滤器 ， 有 =》 52 、 没有 =》 51 
51、静态文件处理 
52、执行过滤器AfterStatic 
53、查找固定路由、正则、自动等匹配 
6、执行controller逻辑 
7、过滤器before exec 
8、执行controller init 、 enable XSRF 是否启动跨域 一般都是 beego.Controller 的初始化，不建议用户继承的时候修改该函数 
9、执行controller prepare 一般用户参数初始化 
10、执行controller get/post 
11、执行controller finish 预留给用户用来重写的，用于释放一些资源。释放在 Init 中初始化的信息数据。 
12、执行过滤器 after exec 13、执行controller destructor

是否开启admin 监控统计url

```