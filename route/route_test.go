package main

import (
	"log"
	"net"
	"net/rpc"
	"testing"

	"github.com/Leon2012/xchat2/protocol"
)

var addr string = "0.0.0.0:7010"

func TestRPC(t *testing.T) {
	address, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		panic(err)
	}
	conn, _ := net.DialTCP("tcp", nil, address)
	defer conn.Close()

	client := rpc.NewClient(conn)
	defer client.Close()

	args := &protocol.NoArg{}
	reply := protocol.NoReply{}
	err = client.Call("RouteRPC.Ping", args, &reply)
	if err != nil {
		log.Fatal("RouteRPC error:", err)
	}
	log.Println(reply)
}
