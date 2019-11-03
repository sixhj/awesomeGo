package main

import (
	"awesomeGo/src/demo_grpc/message"
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:8888")
	if err != nil {
		panic(err.Error()) //nic所在函数F内如果存在要执行的defer函数列表，按照defer的逆序执行
	}

	req := message.InfoReq{
		Name: "six",
		Age:  11,
	}

	funcName1(client, req)
	funcName2(client, req)

	sync(client)

}

func funcName1(client *rpc.Client, req message.InfoReq) {
	var res *message.InfoRes
	client.Call("MathUtil.TestInfo", req, &res)
	fmt.Println(res)                             //name:"seven" address:"zz"
	fmt.Println("&res", &res)                    //  0xc00009e510
	fmt.Println(res.GetName(), res.GetAddress()) //seven zz

}

func funcName2(client *rpc.Client, req message.InfoReq) {
	var res message.InfoRes
	client.Call("MathUtil.TestInfo", req, &res)
	fmt.Println(res)                             //{seven zz {} [] 0}
	fmt.Println("&res", &res)                    // name:"seven" address:"zz"
	fmt.Println(res.GetName(), res.GetAddress()) //seven zz

}

func async(client *rpc.Client) {
	req := 1
	var res *int
	//demoClient.Call("MathUtil.Count", req, &res)  // sync 同步
	result := client.Go("MathUtil.Count", req, &res, nil)
	// async
	r := <-result.Done
	//从返回的channel 中读取变量，读取不到变量就会阻塞，
	fmt.Println(r)
	fmt.Println("返回的参数:", *res)
}

func sync(client *rpc.Client) {
	req := 1
	var res *int
	client.Call("MathUtil.Count", req, &res) // 调用的服务名称  对象.方法名  call 为同步调用
	fmt.Println("返回的参数:", res)               //0xc00009a4b8
	fmt.Println("返回的参数:", *res)              //2
	fmt.Println("返回的参数:", *(&res))           //0xc00009a4b8
	fmt.Println("返回的参数:", &res)              //0xc000010070
	fmt.Println("返回的参数:", &(*res))           //0xc00009a4b8
}
