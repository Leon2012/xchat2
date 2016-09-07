package main

import (
	"net"
	"net/rpc"

	"github.com/Leon2012/xchat2/protocol"
)

type RouteRPC struct {
	store *sessionStore
}

func NewRouteRPC() *RouteRPC {
	r := new(RouteRPC)
	r.store = newSessionStore()
	return r
}

func (r *RouteRPC) Ping(args *protocol.NoArg, reply *protocol.NoReply) error {
	return nil
}

func initRPC() error {
	c := NewRouteRPC()
	rpc.Register(c)
	lis, err := net.Listen("tcp", cfg.Server.Addr)
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
