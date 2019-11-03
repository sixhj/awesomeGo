package main

import (
	pd "awesomeGo/src/demo_grpc/message"
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
)

/* 2019-10-31 liu
Protocol Buffer

syntax = "proto3" 指定版本

message demo {
  string name = 1;   // 类型 名称  =  编号
	// 每个字段都有唯一的编号，一旦消息类型被使用就不应该被更改
  int age = 2;
}

singular：可以有零个或其中一个字段（但不超过一个）
repeated：该字段可以重复任意次数（包括0次），重复值的顺序，将被保留

proto3 中 ，可扩展的repeated字段为数字类型的默认编码

对于go，编译器生成一个.pb.go 文件，其中包含文件中每种消息类型的类型


protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/addressbook.proto



*/
func main() {
	fmt.Println("grpc")

	req := &pd.InfoReq{
		Name: "six",
		Age:  0,
	}

	data, err := proto.Marshal(req) // 序列化
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	fmt.Println(string(data))

	infoReq := &pd.InfoReq{}

	proto.Unmarshal(data, infoReq)

	fmt.Println("反序列化", infoReq.GetName(), infoReq.GetAge())

}

/* 2019-10-31 liu
动态代理技术
client stub 和server stub程序，在具体的编码和实践过程中，都是使用动态代理技术自动生成的一段程序


序列化、反序列化

protobuf协议


*/
