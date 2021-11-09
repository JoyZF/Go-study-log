package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const shortDuration = 2 * time.Second


//func main()  {
//	ctx, cancel := context.WithTimeout(context.Background(), shortDuration)
//	defer cancel()
//	select {
//	case <-time.After(1 * time.Second):
//		fmt.Println("time out")
//	case <- ctx.Done():
//		fmt.Println(ctx.Err())
//	}
//}

func main()  {
	req, err := http.NewRequest("GET", "https://baidu.com", nil)
	if err != nil {
		fmt.Println("http.NewRequest err : %+v",err)
		return
	}
	//将context 作为参数参入 可以做超时控制
	ctx, cancel := context.WithTimeout(req.Context(), 1*time.Second)
	defer cancel()
	req = req.WithContext(ctx)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("http default Client Do err %+v",err)
		return
	}
	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bytes))
}

//函数调用链必须传播上下文
//我们把 context 作为方法首位，本质目的是为了传播context ，自行完整调用链路上的各类控制。
//func list(ctx context.Context,db *sql.DB)   ([]User, error) {
//	ctx, span := trace.StartSpan(ctx, "internal.user.List")
//	defer span.End()
//
//	users := []User{}
//	const q = `SELECT * FROM users`
//
//	if err := db.SelectContext(ctx, &users, q); err != nil {
//		return nil, errors.Wrap(err, "selecting users")
//	}
//
//	return users, nil
//}

//context的继承和派生
func handle(w http.ResponseWriter,req *http.Request)  {
	//parent context
	timeout, _ := time.ParseDuration(req.FormValue("timeout"))
	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	//children context
	newCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	newCtx.Done()
}

//一般会有父级context 和子级context 的区别，我们要保证程序的行为中上下文对于多个goroutine同时使用是安全的。
//并且存在父子级别关系，父级context关闭或超时，可以影响到子级context的程序。


//但在实际的 context 建议中，我们会建议使用 context.TODO 方法来创建顶级的 context，直到弄清楚实际 Context 的下一步用途，再进行变更。
//总结
//
//对第三方调用要传入 context，用于控制远程调用。
//不要将上下文存储在结构类型中，尽可能的作为函数第一位形参传入。
//函数调用链必须传播上下文，实现完整链路上的控制。
//context 的继承和派生，保证父、子级 context 的联动。
//不传递 nil context，不确定的 context 应当使用 TODO。
//context 仅传递必要的值，不要让可选参数揉在一起。