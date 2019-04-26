package main

import (
	"net/rpc"
	"gorpc/jsonServer/impl"
	"net"
	"fmt"
	"os"
	"net/rpc/jsonrpc"
)

func main() {
	//注册服务
	rpc.Register(new(impl.Airth))
	lis,err:=net.Listen("tcp",":8000")
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Fprintln(os.Stdout,"start connection")
	for  {
		conn,errs := lis.Accept()// 接收客户端连接请求
		if errs!=nil {
			fmt.Fprintln(os.Stderr,errs)
			continue
		}
		go func(c net.Conn) {
			fmt.Fprintln(os.Stdout,"new client in coming\n")
			jsonrpc.ServeConn(c)
		}(conn)
	}
}