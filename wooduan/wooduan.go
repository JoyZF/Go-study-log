package wooduan

import (
	"context"
)

type Args struct {
	A int
	B int
}

type Reply struct {
	C int
}

type Arith int
//Mul
// 1、可导出类型的方法
// 2、 接受3个参数，第一个是context 其他两个都是可导出的
// 3、 第三个参数是指针
// 4、 有一个error类型的返回值
func (a *Arith) Mul(ctx context.Context, args *Args,reply *Reply) error {
	reply.C = args.A * args.B
	return nil
}






