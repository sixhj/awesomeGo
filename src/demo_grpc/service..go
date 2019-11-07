package main

import (
	"awesomeGo/src/demo_grpc/message"
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

/* 2019-10-31 liu
rpc 包的使用
*/
func main() {

	mathUtil := new(MathUtil)
	// 注册服务
	register := rpc.Register(mathUtil) // 默认的服务名就是 方法对象的名称

	//rpc.RegisterName("服务名称"，mathUtil) // 可以指定服务名称
	if register != nil {
		panic(register.Error())
	}
	// 将服务注册到http协议上，调用者就可以调用http的方式进行数据传递
	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println("service  run ")
	http.Serve(listener, nil)

}

// 定义一个对象
type MathUtil struct {
}

// 定义一个对象的方法  req 请求参数   res 响应指针类型的返回值
func (m *MathUtil) Count(req int, res *int) error {
	// 服务
	fmt.Println("服务被调用")
	add = req + req
	res = &add // 返回的是一个地址
	return nil
}

// 如果需要传递多个参数的话可以使用结构体
func (m *MathUtil) TestInfo(req message.InfoReq, res *message.InfoRes) error {
	unix := time.Now().Unix()
	fmt.Println(unix)

	fmt.Println("name:", req.GetName(), "age:", req.GetAge())

	infoRes := message.InfoRes{
		Name:    "seven",
		Address: "zz",
	}
	fmt.Println("infoRes:", infoRes)   //infoRes: {seven zz {} [] 0}
	fmt.Println("&infoRes:", &infoRes) //name:"seven" address:"zz"
	//fmt.Println("*infoRes:",*infoRes) //  报错

	*res = infoRes
	fmt.Println("res:", res)         //res: name:"seven" address:"zz"
	fmt.Println("*res:", *res)       //{seven zz {} [] 0}
	fmt.Println("&res:", &res)       // 0xc000010078
	fmt.Println("&(*res):", &(*res)) //name:"seven" address:"zz"
	//fmt.Println("&(*res):",&(&res)) //报错
	//fmt.Println("&(*res):",*(*res)) //报错

	return nil
}
