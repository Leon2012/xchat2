package main

import (
	"github.com/Leon2012/xchat2/protocol"
)

type Account struct {
}

func (model *Account) Login(args *protocol.LoginArg, reply *protocol.LoginReply) error {
	return nil
}
