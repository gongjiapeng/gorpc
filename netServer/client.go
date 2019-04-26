package main

import (
	"net/rpc"
	"fmt"
	"gorpc/netServer/impl"
)
func main() {
	conn,err:=rpc.DialHTTP("tcp",":8000")
	if err!=nil {
		fmt.Println(err)
	}

	req := impl.AirthReq{A:9,B:10}
	var res impl.AirthResp
	//乘法
	errs := conn.Call("Airth.Multiply",req,&res)
	if errs!=nil {
		fmt.Println(errs)
	}
	fmt.Println(res.Pro)
}
