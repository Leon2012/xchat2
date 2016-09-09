package main

import (
	"net"
	"net/rpc"
)

func initRPC() error {
	account := &Account{}
	topic := &Topic{}
	rpc.Register(account)
	rpc.Register(topic)
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
