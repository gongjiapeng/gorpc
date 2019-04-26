/*
@Time : 19/4/26 下午4:46
@Author : gongjiapeng
@File : impl
@Software: GoLand
*/
package impl

import "errors"

type Airth struct {
}
//请求参数结构体
type AirthReq struct {
	A int
	B int
}
//响应参数结构体
type AirthResp struct {
	Pro int
	Quo int
	Rem int
}
//乘法运算
func (t *Airth) Multiply(req *AirthReq,resp *AirthResp) error {
	resp.Pro = req.A * req.B
	return nil
}
//除法运算
func (t *Airth) Divide(req *AirthReq,resp *AirthResp) error {
	if req.B==0 {
		return errors.New("除数不能为0")
	}
	resp.Rem = req.A % req.B
	resp.Quo = req.A / req.B
	return nil
}

