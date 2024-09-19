package main

import (
	"fmt"
	"net/rpc"
)

// 定义参数结构
type Args struct {
	A, B int
}

func main() {
	// 连接RPC服务器
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("连接服务器失败：", err)
		return
	}
	defer client.Close()

	// 定义参数
	args := &Args{7, 8}
	var reply int

	// 调用远程方法
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		fmt.Println("调用失败：", err)
		return
	}

	// 输出结果
	fmt.Println("乘法结果：", reply)
}
