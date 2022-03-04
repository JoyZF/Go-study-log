package main

import (
	"bytes"
	gateway "github.com/rpcxio/rpcx-gateway"
	"github.com/smallnest/rpcx/codec"
	"io/ioutil"
	"log"
	"net/http"
)

type Args struct {
	A int
	B int
}

type Reply struct {
	C int
}

func main() {
	cc := codec.MsgpackCodec{}
	args := &Args{
		A: 10,
		B: 20,
	}

	data, _ := cc.Encode(args)

	req, err := http.NewRequest("POST", "http://127.0.01:9981/", bytes.NewReader(data))
	if err != nil {
		log.Fatal("failed to create request:", err)
		return
	}

	// set extra headers
	h := req.Header
	h.Set(gateway.XMessageID, "10000")
	h.Set(gateway.XMessageType, "0")
	h.Set(gateway.XSerializeType, "3")
	h.Set(gateway.XServicePath, "Arith")
	h.Set(gateway.XServiceMethod, "Mul")

	// send to gateway
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("failed to call: ", err)
	}
	defer res.Body.Close()

	// handle http response
	replyData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("failed to read response: ", err)
	}

	// parse reply
	reply := &Reply{}
	err = cc.Decode(replyData, reply)
	if err != nil {
		log.Fatal("failed to decode reply: ", err)
	}

	log.Printf("%d * %d = %d", args.A, args.B, reply.C)

}