package main

import (
	"fmt"
	"math/rand"
	"time"
)

/* 2019-09-20 liu
同步

使用channel

通道通过解决数据共享问题，让并发编程变得更加清晰。
通道是一个通信管道，它用于go协程之间传递数据。
换句话说，go协程可以通过通道，传递数据给另外一个go协程。
其结果就是，在任何时候，仅有一个go协程可以访问数据。
通道与所有其他的东西一样，也有类型。这个类型，就是将要在通道中传递的数据的类型。


从一个通道接收或者发送数据时会阻塞。
当我们从一个通道接收数据时，直到数据可用, go协程才会继续执行。
类似的，往一个通道发送数据时，在数据被接收之前, go协程也不会继续执行。

go语言确保，往一个通道发送数据时，仅有一个单独的接收器可以接收。


通道是唯一共享的状态，通过通道，可以安全的，并发发送和接收数据。
通道提供了我们需要的所有同步代码，并且也确保，在任意的特定时刻，

只有一个go协程，可以访问数据的特定部分。



*/

func main() {
	c := make(chan int)

	//CHANNEL <- DATA  //往通道发送数据
	//VAR := <-CHANNEL  //从通道接收数据

	//使用select的最主要目的是，通过它管理多个通道。
	// 给定多个通道，select将阻塞直到有一个通道可用。
	// 如果没有可用的通道，当提供了default语句时，执行该分支。
	// 当多个通道都可用时，选择其中的一个通道是随机的。
	for {
		select {
		case c <- rand.Int():
			//可选的代码
		default:
			//这里可以留下空行以丢弃数据
			fmt.Println("dropped")
		}
		time.Sleep(time.Millisecond * 50)
	}
}

/* 2019-09-21 liu
# 6.3 通道

并发编程的挑战主要在于数据共享。如果你的go协程没有共享数据，你就不需要担心同步他们。但是，对于所有的系统，这不是一个选择。实际上，很多系统以完全相反的目标构建：在多个请求中共享数据。内存缓存或者数据库都是这方面的好例子。这正变得越来越流行的事实。

通道通过解决数据共享问题，让并发编程变得更加清晰。通道是一个通信管道，它用于go协程之间传递数据。换句话说，go协程可以通过通道，传递数据给另外一个go协程。其结果就是，在任何时候，仅有一个go协程可以访问数据。

通道与所有其他的东西一样，也有类型。这个类型，就是将要在通道中传递的数据的类型。例如，创建一个通道，这个通道可以用来传递一个整数，我们可以这样：


    c := make(chan int)


这个通道的类型是`chan int`。因此，将这个通道传递给一个函数，可以这样声明：

    func worker(c chan int) { ... }

通道支持2种操作：接收和发送。我们可以使用下面方式往通道发送数据：

    CHANNEL <- DATA

可以使用下面方式从通道接收数据：

    VAR := <-CHANNEL

箭头的方向就是数据的流动方向。当发送数据时，数据流入通道。当接收数据时，数据流出通道。

最后，在查看我们的第一个例子之前，我们需要知道从一个通道接收或者发送数据时会阻塞。当我们从一个通道接收数据时，直到数据可用, go协程才会继续执行。类似的，往一个通道发送数据时，在数据被接收之前, go协程也不会继续执行。

假设这种情况：对输入数据，我们想通过不同的协程去处理。这是一种常见的需求。如果通过go协程接收输入的数据，并进行数据密集型处理，那么,客户端会有超时风险。首先，我们将写出`worker`。这可以是一个简单的函数，但是我会让它变成一个结构体的部分，因为我们之前从来没有这样使用过go协程：

```go
type Worker struct {
    id int
}

func (w Worker) process(c chan int) {
    for {
        data := <-c
        fmt.Printf("worker %d got %d\n", w.id, data)
    }
}
```

我们的`worker`很简单。它会一直等待数据，直到数据可用, 然后处理它。它在一个循环中，永远尽职的等待更多的数据并处理。

为了使用上面的`worker`，我们首先要做的是启动一些`worker`：

```go
c := make(chan int)
for i := 0; i < 4; i++ {
    worker := Worker{id: i}
    go worker.process(c)
}
```

然后我们可以给它们一些工作：

```go
for {
    c <- rand.Int()
    time.Sleep(time.Millisecond * 50)
}
```

这是完整的代码，运行它：

```go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    c := make(chan int)
    for i := 0; i < 5; i++ {
        worker := &Worker{id: i}
        go worker.process(c)
    }

    for {
        c <- rand.Int()
        time.Sleep(time.Millisecond * 50)
    }
}

type Worker struct {
    id int
}

func (w *Worker) process(c chan int) {
    for {
        data := <-c
        fmt.Printf("worker %d got %d\n", w.id, data)
    }
}
```

我们不知道哪个`worker`将获得数据。我们所知道的是，go语言确保，往一个通道发送数据时，仅有一个单独的接收器可以接收。

注意：通道是唯一共享的状态，通过通道，可以安全的，并发发送和接收数据。通道提供了我们需要的所有同步代码，并且也确保，在任意的特定时刻，只有一个go协程，可以访问数据的特定部分。

## 6.3.1 带缓存的通道

在上面的代码中，如果输入的数据，超过我们的处理能力，会发生什么？你可以模拟这种场景，在`worker`接收到数据后，让`worker`执行`time.Sleep`：

```go
for {
    data := <-c
    fmt.Printf("worker %d got %d\n", w.id, data)
    time.Sleep(time.Millisecond * 500)
}
```

在`main`函数中会发什么呢？接收用户的输入数据（这里通过一个随机的数字生成器模拟）会被阻塞，因为往通道发送数据时，没有可用的接收者。

在这种情况下，你需要确保数据被处理，你可能想要让客户端阻塞。在其他情况下，你可能愿意不确保数据被处理。这里有一些流行的策略能完成此事。首先是将数据缓存起来。如果没有`worker`可用，我们想将数据，暂时存放在一个有序的队列中。通道实现了这种缓存功能。当我们使用`make`创建一个通道时，我们可以指定通道的长度：

    c := make(chan int, 100)

你可以这样调整，但是你将注意到，处理过程仍然不顺利。带缓存通道没有提供更多的功能；它只不过是为挂起的作业提供一个队列，以一种更好的方式处理数据突然飙升。在我们示例中，我们可以连续不断的发送更多的、超出`worker`处理能力的数据。

然而，通过查看通道的长度，我们可以了解到，带缓存通道中有待处理的缓存数据：

```go
for {
    c <- rand.Int()
    fmt.Println(len(c))
    time.Sleep(time.Millisecond * 50)
}
```

你可以看到，带缓存通道的长度在不断增长，直到装满为止，到时，往通道发送的数据又开始被阻塞。

## 6.3.2 select

即使借助缓存，我们还是需要开始丢弃一些消息。我们不能使用一个无限大的内存，并指望人工的释放它。所以我们使用go语言的`select`。

在语法结构上，`select`看起来有点类似`switch`。通过`select`，我们能写出解决通道不可写问题的代码。首先，让我们去掉通道的缓存，这样可以更清晰的看到`select`是如何工作的。

    c := make(chan int)

接下来，我们修改`for`循环：

```go
for {
    select {
        case c <- rand.Int():
            //可选的代码
        default:
            //这里可以留下空行以丢弃数据
            fmt.Println("dropped")
    }
    time.Sleep(time.Millisecond * 50)
}
```

我们每秒往通道中发送20个信息，但是我们的程序，每秒只能处理10个信息；因此，有一半的信息被丢弃。

这仅仅只是我们使用`select`完成一些事的开始。使用`select`的最主要目的是，通过它管理多个通道。给定多个通道，`select`将阻塞直到有一个通道可用。如果没有可用的通道，当提供了`default`语句时，执行该分支。当多个通道都可用时，选择其中的一个通道是随机的。

很难想出一个简单的例子来证明这种行为，因为这是一种高级特性。在下一小节可能有助于说明这个问题。

## 6.3.3 超时

我们已经学习了缓存信息，并且丢弃它们的简单做法。另外一种比较流行的做法是使用超时。我们将阻塞一段时间，但是不是一直阻塞。这在go中很容易实现。老实说，这个语法有点难接受，确是非常灵活、有用的特性，我不能不介绍它。

为了使阻塞达到最大值，我们可以使用`time.After`函数。让我们看看它会发生什么神奇的事。为了使用这种方式，我们数据发送变成：

```go
for {
    select {
        case c <- rand.Int():
        case <-time.After(time.Millisecond * 100):
            fmt.Println("timed out")
    }
    time.Sleep(time.Millisecond * 50)
}
```

`time.After`将返回一个通道，所以我们可以对它使用`select`语句。这个通道在经过指定的时间后会被写入。就是这样。没有什么比这个更神奇了。如果你依然觉得奇怪，这里实现了一个`after`，如下所示：

```go
func after(d time.Duration) chan bool {
    c := make(chan bool)
    go func() {
        time.Sleep(d)
        c <- true
    }()
    return c
}
```

回到我们的`select`语句，这里有一些好玩的东西。首先，如果你在后面添加`default`分之会发生什么？你能猜到吗？试试。如果你不确定会发生什么，记住如果通道不可用的话，`default`分支会被立即执行。

此外，`time.After`是一个`chan time.Time`类型的通道。在上面的例子中，我们将发送给通道的值简单丢弃掉。如果你想，你也可以获取到这个值：

```go
    case t := <-time.After(time.Millisecond * 100):
        fmt.Println("timed out at", t)
```

密切注意我们的`select`。注意我们正在往`c`中发送数据，但是从`time.After`收取。不管我们是从通道中接收数据、发送数据或者收发数据，`select`工作机制都一样：

- 第一个可用的通道被选中
- 如果多个通道可用，随机选中一个通道
- 如果没有通道可用，`default`分之被执行
- 如果没有`default`分支，`select`将阻塞

最后，在`for`循环中使用`select`也是比较常见的，例如：

```go
for {
    select {
        case data := <-c:
            fmt.Printf("worker %d got %d\n", w.id, data)
        case <-time.After(time.Millisecond * 10):
            fmt.Println("Break time")
            time.Sleep(time.Second)
    }
}
```

*/
