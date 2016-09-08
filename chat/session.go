package main

import (
	"time"

	"github.com/gorilla/websocket"
)

const (
	NONE = iota
	WEBSOCK
	TCP
)

type Session struct {
	proto       int
	ws          *websocket.Conn
	remoteAddr  string
	userAgent   string
	ver         int
	deviceId    string
	uid         int64
	lastTouched time.Time
	lastAction  time.Time

	send   chan []byte
	stop   chan []byte
	detach chan string

	// Session ID
	sid string
}
