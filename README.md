go语言
https://www.practical-go-lessons.com
Mastering Go 这本书很好

https://github.com/talkgo/read

GMP:
https://morsmachine.dk/go-scheduler
https://rakyll.org/scheduler/
https://medium.com/the-polyglot-programmer/what-are-goroutines-and-how-do-they-actually-work-f2a734f6f991
https://blog.csdn.net/u010853261/article/details/84790392
https://hustcat.github.io/dive-into-goroutine/
https://tiancaiamao.gitbooks.io/go-internals/content/zh/05.1.html
https://calvinfeng.gitbook.io/gonotebook/concurrency/04-02-go-concurrency-part-1
https://medium.com/random-life-journal/scheduling-in-go-727c9b88c93a

Golang修养之路 刘丹冰
https://www.kancloud.cn/aceld/golang/1858955
https://www.kancloud.cn/aceld/golang/1958305
https://www.bilibili.com/video/BV19r4y1w7Nx?p=6&spm_id_from=pageDriver

Effective Go 中文文档   https://gocn.vip/wiki/effective
Network programming with Go	https://jan.newmarch.name/go/
深入理解Go语言的Channels特性	https://www.s0nnet.com/archives/go-channels-behavior
Go语言之Channels实际应用	https://www.s0nnet.com/archives/go-channels-practice
深入解析Go	https://tiancaiamao.gitbooks.io/go-internals/content/zh/
图解Go的channel底层原理
https://mp.weixin.qq.com/s/40uxAPdubIk0lU321LmfRg
go101 https://go101.org/article/101.html
About Go 101	https://go101.org/article/101-about.html
Gopher 2019 Go并发编程的分享	https://colobu.com/2019/04/28/gopher-2019-concurrent-in-action/
golang channel 使用总结	http://litang.me/post/golang-channel/
《The Go Programming Language》
https://beyondkmp.com/books/golang/The.Go.Programming.Language.pdf
Go语言高级编程(Advanced Go Programming)	
https://chai2010.cn/advanced-go-programming-book/
曹大 Go语言高级编程 https://chai2010.cn/advanced-go-programming-book/
深入Go的底层，带你走近一群有追求的人
https://mp.weixin.qq.com/s/obnnVkO2EiFnuXk_AIDHWw
曹大博客 https://xargin.com/
曹大go源码分析 https://github.com/cch123/golang-notes
曹大汇编视频 https://www.bilibili.com/video/av46494102
https://github.com/cch123/asmshare/blob/master/layout.md
https://github.com/talkgo/night/issues/186
https://github.com/golang/go/wiki
Go夜读	https://talkgo.org/

https://golangbot.com/learn-golang-series/

https://golang.org/dl/

https://datatracker.ietf.org/doc/html/rfc7231#section-6.1

gin:
https://gin-gonic.com/
https://github.com/gin-gonic/gin

defer:
https://segmentfault.com/a/1190000018169295#articleHeader4

gorush：Go 编写的通知推送服务器。
fnproject：容器原生，云 serverless 平台。
photoprism：基于 Go 和 Google TensorFlow 实现的个人照片管理工具。
krakend：拥有中间件的超高性能 API 网关。
picfit：Go 编写的图像尺寸调整服务器。
gotify：基于 WebSocket 进行实时消息收发的简单服务器。
cds：企业级持续交付和 DevOps 自动化开源平台。

云原生:
高内聚 低耦合 自动化 平台化

云原生的本质其实是基础设施与业务的解耦，以及基础设施自身的标准化。

早先的 k8s 的诞生一定程度上抹平了二线公司与一线公司的基础设施差距，以低成本的方式让每家公司都拥有了曾经想都不敢想的故障自愈和自动扩缩容功能。


[mattson2004patterns] Mattson, Timothy G, Beverly Sanders, and Berna Massingill. 2004. Patterns for Parallel Programming. Pearson Education.



Comparison Operators

Operator Signification
== equal
!= not equal
> greater
>= greater or equal
< less
<= less or equal





网络编程:


下面是具体的过程：
客户端的协议栈向服务器端发送了 SYN 包，并告诉服务器端当前发送序列号 j，客户端进入 SYNC_SENT 状态；
服务器端的协议栈收到这个包之后，和客户端进行 ACK 应答，应答的值为 j+1，表示对 SYN 包 j 的确认，同时服务器也发送一个 SYN 包，告诉客户端当前我的发送序列号为 k，服务器端进入 SYNC_RCVD 状态；
客户端协议栈收到 ACK 之后，使得应用程序从 connect 调用返回，表示客户端到服务器端的单向连接建立成功，客户端的状态为 ESTABLISHED，同时客户端协议栈也会对服务器端的 SYN 包进行应答，应答数据为 k+1；
应答包到达服务器端后，服务器端协议栈使得 accept 阻塞调用返回，这个时候服务器端到客户端的单向连接也建立成功，服务器端也进入 ESTABLISHED 状态。

为什么tcp建立连接需要三次握手，解释如下:
tcp连接的双方要确保各自的收发消息的能力都是正常的。
客户端第一次发送握手消息到服务端，
服务端接收到握手消息后把ack和自己的syn一同发送给客户端，这是第二次握手，
当客户端接收到服务端发送来的第二次握手消息后，客户端可以确认“服务端的收发能力OK，客户端的收发能力OK”，但是服务端只能确认“客户端的发送OK，服务端的接收OK”，
所以还需要第三次握手，客户端收到服务端的第二次握手消息后，发起第三次握手消息，服务端收到客户端发送的第三次握手消息后，就能够确定“服务端的发送OK，客户端的接收OK”，
至此，客户端和服务端都能够确认自己和对方的收发能力OK，tcp连接建立完成。

这个问题的本质是, 信道不可靠, 但是通信双发需要就某个问题达成一致. 而要解决这个问题, 无论你在消息中包含什么信息, 三次通信是理论上的最小值. 所以三次握手不是TCP本身的要求, 而是为了满足"在不可靠信道上可靠地传输信息"这一需求所导致的。

关于三次握手
1.信道不安全 保证通信需要一来一回
2.客户端的来回和服务端的来回 共四次 这是最多四次
3.客户端的回和服务端的来合并成一个，就是那个sync k ack j+1
4.这样就是三次握手

非阻塞调用的场景就是高性能服务器编程！我所有的调用都不需要等待对方准备好了再返回，而是立即返回，那么我怎么知道是否准备好了？就是把这些fd注册到类似select或者epoll这样的调用中，变多个fd阻塞为一个fd阻塞，只要有任何一个fd准备好了，select或者epoll都会返回，然后我们在从中取出准备好了的fd进行各种IO操作。

golang学习笔记:
根本是这本书:www.practical-go-lessons.com
学习一门程序设计语言，最基本的内容:
变量
var roomNumber int





Warning : short variable declaration cannot be used outside functions !

constant:
const version string = "1.3.2"
const version = "1.3.2"

指针
var p *int
var answer int = 42
p = &answer
fmt.Println(p)
// 0xc00000a098

Dereferencing/referencing
With the operator * you follow the address
With the operator & you take the address


数组
var myArray [2]int
    myArray[0] = 156
    myArray[1] = 147
    fmt.Println(myArray)

OR
myArray := [2]int{156,147}
myArray := [...]int{156, 147}

二维数组:
var myValue [3][2]int

map
var m = make(map[int]int)
cities := make(map[string]string)
for k, v := range employeeMap {
    fmt.Printf("Key: %s - Value: %s\n", k, v)
}

if value, ok := m[n]; ok {
    return value
}


To add a pair composed of a key and a element simply use the following syntax :
m[key]=value




The key is passed to the hash function.The hash function will return the hash.
From the hash, we will retrieve the bucket id.
Go will then iterate over the bucket elements to find a place to store the key and the element.
When the key is already present, Go will override the element’s value.

Map values are not addressable.
Why this behavior? Because Go can change the memory location of a key-value pair when it adds a new key-value pair. Go will do this under the hood to keep the complexity of retrieving a key-value pair at a constant level. As a consequence, the address can become invalid. Go prefers to forbid the access of a possible invalid address than letting you try your chance. This is a good thing !

分片
test := make([]int,2,10)

s := make([]int, 3)
   s[0] = 12
   s[2] = 3

s := e[low:high]
high 是slice的长度，不是最高位的坐标

func multiply(slice []int, factor int) {
    for i := 0; i < len(slice); i++ {
        slice[i] = slice[i] * factor
    }
}

names := []string{"John", "Bob", "Claire", "Nik"}
for i, name := range names {
    fmt.Println("Element at index", i, "=", name)
}

names := []string{"John", "Bob", "Claire", "Nik"}
for _, name := range names {
    name = strings.ToUpper(name)
}
fmt.Println(names)


struct
type Cat struct {
  	Color string
	Age uint8
 	Name string
    }

type ListNode struct {
   Val int
   Next *ListNode
}

link1 := &ListNode{}
link1.Val = 1
link2 := &ListNode{}
link2.Val = 2
link3 := &ListNode{}
link3.Val = 3
link4 := &ListNode{}
link4.Val = 4
link5 := &ListNode{}
link5.Val = 5

link1.Next = link2
link2.Next = link3
link3.Next = link4
link4.Next = link5


for
 myArray := [2]int{156, 147}
   for index, element := range myArray {
       fmt.Printf("%d is %d\n", index, element)
   }

for _, element := range myArray {
    fmt.Println(element)
   }

for i := 0; i < len(a); i++ {
      fmt.Println(i, a[i])
    }



function


func computePrice(rate float32, nights int) float32 {
    return rate * float32(nights)
}


func computePrice(rate float32, nights int) (price float32) {
   price = rate * float32(nights)
   return
}


func vacantRooms() int {
    rand.Seed(time.Now().UTC().UnixNano())
    return rand.Intn(100)
}

func printHeader() {
    fmt.Println("Hotel Golang")
    fmt.Println("San Francisco, CA")
}

package main

func testa(x int) (func(), func()) {
   return func() {
         println(x)
         x += 10
      }, func() {
         println(x)
      }
}
func main() {
   a, b := testa(100)
   a()
   b()
}



method


type Item struct {
     ID string
   }

type Cart struct {
    ID        string
    CreatedAt time.Time
    UpdatedAt time.Time
    lockedAt  time.Time
    user.User
    Items        []Item
    CurrencyCode string
    isLocked     bool
}

func (c *Cart) TotalPrice() (*money.Money, error) {
    // ...
    return nil, nil
}

func (c *Cart) Lock() error {
    // ...
    return nil
}


method 存在的意义是为 struct 赋予能力，receiver 就是 struct.
interface 是为了支持多态.

世界上没有高深的技能，简单的动作重复练习就是高深的技能。

channel
ch1 := make(chan int)
ch2 := make(chan string, 3)
x, ok = <-ch 
if !ok {
    log.Println("channel is empty or closed")
}


https://golang.design/research/ultimate-channel/
https://morsmachine.dk/go-scheduler
http://www.minaandrawos.com/2015/12/06/concurrency-in-golang/

Go relies on a concurrency model called CSP ( Communicating Sequential Processes) 

How to remember the syntax easily ? 
When the arrow points to the channel we send data into the channel，
Otherwise we receive data from the channel

package main

import "fmt"

func main() {
    ch5 := make(chan int, 2)
    ch5 <- 42
    close(ch5)
    received, ok := <-ch5
    fmt.Println(received, ok)
}


This program will output :
42 true


package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3) //带有缓冲的channel

	fmt.Println("len(c) = ", len(c), ", cap(c)", cap(c))

	go func() {
		defer fmt.Println("子go程结束")

		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("子go程正在运行, 发送的元素=", i, " len(c)=", len(c), ", cap(c)=", cap(c))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 4; i++ {
		num := <-c //从c中接收数据，并赋值给num
		fmt.Println("num = ", num)
	}

	fmt.Println("main 结束")
}



package main

import (
   "fmt"
   "sync"
   "time"
)

func main() {
   fmt.Printf("Program start \n")
   // initialize the wait group
   var waitGroup sync.WaitGroup
   waitGroup.Add(10)
   for i := 0; i < 10; i++ {
      go concurrentTaks(i, &waitGroup)
   }
   waitGroup.Wait()
   finishTask()
   fmt.Printf("Program end \n")

}

func finishTask() {
   fmt.Println("Executing finish task")
}

func concurrentTaks(taskNumber int, waitGroup *sync.WaitGroup) {
   fmt.Printf("BEGIN Execute task number %d \n", taskNumber)
   time.Sleep(100 * time.Millisecond)
   fmt.Printf("END Execute task number %d \n", taskNumber)
   waitGroup.Done()
}


曹春晖 Go语言高级编程 https://chai2010.cn/advanced-go-programming-book/
