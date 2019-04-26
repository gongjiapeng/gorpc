/*
@Time : 19/4/26 下午3:29 
@Author : gongjiapeng
@File : server.go
@Software: GoLand
*/
package main

import (
	"net/rpc"
	"net"
	"fmt"
	"net/http"
	"gorpc/netServer/impl"
)
func main() {
	//注册服务
	rpc.Register(new(impl.Airth))
	//采用http协议作为rpc载体
	rpc.HandleHTTP()

	lis ,err := net.Listen("tcp",":8000")
	if err!=nil {
		fmt.Println(err)
	}
	http.Serve(lis,nil)
}


