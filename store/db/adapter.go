package db

import "github.com/Leon2012/xchat2/store/types"

type Adapter interface {
	Open(c interface{}) error
	Close() error
	IsOpen() bool

	Login(token string) (*types.User, error)
	Logout(*types.User) error
}
