package client

import (
	"net"
	"net/rpc"
)

type RPCClient struct {
	client *rpc.Client
	conn   *net.TCPConn
}

func NewRPCClient() *RPCClient {
	return &RPCClient{}
}

func (c *RPCClient) Open(addr string) error {
	address, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return err
	}
	conn, err := net.DialTCP("tcp", nil, address)
	if err != nil {
		return err
	}
	c.conn = conn
	client := rpc.NewClient(conn)
	c.client = client
	return nil
}

func (c *RPCClient) Call(method string, args interface{}, reply interface{}) error {
	err := c.client.Call(method, args, reply)
	return err
}

func (c *RPCClient) Close() {
	c.conn.Close()
	c.client.Close()
	c.client = nil
	c.conn = nil
}
