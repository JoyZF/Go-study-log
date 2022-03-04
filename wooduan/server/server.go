package main

import (
	"context"
	"github.com/smallnest/rpcx/server"
	"study/wooduan"
	"time"
)

type Server struct {
	readTimeout time.Duration
	writeTimeout time.Duration
}

func mul(ctx context.Context, args *wooduan.Args, reply *wooduan.Reply) error {
	reply.C = args.A * args.B
	time.Sleep(3 * time.Second)
	return nil
}


func main()  {
	s := server.NewServer(
		server.WithReadTimeout(1 * time.Second),
		server.WithWriteTimeout(1 * time.Second),
		)
	_ = s.RegisterName("Arith", new(wooduan.Arith), "")
	s.RegisterFunction("a.fake.service", mul, "")
	_ = s.Serve("tcp", ":8972")
}