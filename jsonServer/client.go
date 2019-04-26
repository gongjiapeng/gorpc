package main

import (
	"net/rpc/jsonrpc"
	"fmt"
	"gorpc/jsonServer/impl"
)

func main() {
	conn,err:=jsonrpc.Dial("tcp",":8000")
	if err!=nil {
		fmt.Println(err)
	}
	req:=impl.AirthReq{A:19,B:10}
	var res impl.AirthResp
	errs := conn.Call("Airth.Divide",req,&res)
	if errs!=nil {
		fmt.Println(errs)
	}
	fmt.Printf("%d / %d, quo is %d, rem is %d\n", req.A, req.B, res.Quo, res.Rem)
}