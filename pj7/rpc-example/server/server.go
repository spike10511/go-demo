package main

import (
	"fmt"
	"net"
	"net/rpc"
)

// 定义一个服务
type Arith int

// 定义方法：Multiply
func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

// 定义参数结构
type Args struct {
	A, B int
}

func main() {
	// 创建服务实例
	arith := new(Arith)

	// 注册服务
	rpc.Register(arith)

	// 监听TCP连接
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("监听失败：", err)
		return
	}
	defer l.Close()

	fmt.Println("服务已启动，等待客户端连接...")

	// 接受并处理客户端连接
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("连接接受失败：", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
