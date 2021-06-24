package main

import (
	"log"
	"net"
	"os"
	"study/redis/protocol"
)

const (
	Address = "r-2zefrvlwps27pj1eh1pd.redis.rds.aliyuncs.com:6379?nodename=none&pass=Tianzhuo@2020"
	Network= "tcp"
)

func Conn(network, address string) (net.Conn, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil,err
	}
	return conn,err
}

func main()  {
	args := os.Args[1:]
	if len(args) <= 0 {
		log.Fatalf("Os.Args <= 0 ")
	}

	reqCommand := protocol.GetRequest(args)
	redisConn, err := Conn(Network, Address)
	if err != nil {
		log.Fatalf("Conn err : %v",err)
	}
	_, err = redisConn.Write(reqCommand)
	if err != nil {
		log.Fatalf("Conn write err : %v",err)
	}

	command := make([]byte, 1024)
	n, err := redisConn.Read(command)
	if err != nil {
		log.Fatalf("Conn read err : %v",err)
	}
	reply, err := protocol.GetReply(command[:n])
	if err != nil {
		log.Fatalf("Conn get reply err %v",err)
	}
	log.Println("Reply %v",reply)
	log.Println("Command $v",string(command[:n]))
}


