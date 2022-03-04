package main

import (
	"context"
	"github.com/smallnest/rpcx/client"
	"log"
	"study/wooduan"
	"time"
)

//func main()  {
//	d, _ := client.NewPeer2PeerDiscovery("tcp@127.0.0.1:8972", "")
//
//	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
//	defer xclient.Close()
//
//
//	args := &wooduan.Args{
//		A: 10,
//		B: 20,
//	}
//
//	reply := wooduan.Reply{}
// 同步调用
//	err := xclient.Call(context.Background(), "Mul", args, reply)
//	if err != nil {
//		log.Fatalf("failed to call: %v", err)
//	}
//
//	log.Printf("%d * %d = %d", args.A, args.B, reply.C)
//}

func main() {
	d, _ := client.NewPeer2PeerDiscovery("tcp@127.0.0.1:8972", "")
	xClient := client.NewXClient("a.fake.service", client.Failfast, client.ConsistentHash, d, client.DefaultOption)
	defer xClient.Close()


	args := &wooduan.Args{
		A: 10,
		B: 20,
	}

	for  {
		reply := &wooduan.Reply{}
		// 异步调用
		 err := xClient.Broadcast(context.Background(), "mul", args, reply)
		if err != nil {
			log.Fatalf("failed to call: %v", err)
		}

		log.Printf("%d * %d = %d", args.A, args.B, reply.C)
		time.Sleep(1e9)
	}
}