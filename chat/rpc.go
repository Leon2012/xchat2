package main

import (
	"net"
	"net/rpc"

	"github.com/Leon2012/xchat2/protocol"
)

type ChatRPC struct {
	
}

func NewChatRPC() *ChatRPC {
	r := new(ChatRPC)
	return r
}

func (r *ChatRPC) Ping(args *protocol.NoArg, reply *protocol.NoReply) error {
	return nil
}

func initRPC() error {
	c := NewChatRPC()
	rpc.Register(c)
	lis, err := net.Listen("tcp", cfg.Rpc.Addr)
	if err != nil {
		return err
	}
	go func() {
		for {
			conn, err := lis.Accept()
			if err != nil {
				continue
			}
			go rpc.ServeConn(conn)
		}
	}()
	return nil
}
