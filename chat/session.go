package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Leon2012/xchat2/protocol"
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

	send   chan []byte //send message
	stop   chan []byte
	detach chan string
	subs   map[string]*Subscription // topic name -> Subscription

	// Session ID
	sid string
}

type Subscription struct {
	uaChange chan<- string
}

//out message
func (s *Session) queueOut(msg *protocol.ServerComMessage) {
	if s == nil {
		return
	}
	data, _ := json.Marshal(msg)
	appLogger.Debug("session.queueOut send '%s' ", string(data))
	select {
	case s.send <- data:
	case <-time.After(time.Millisecond * 10): //超时, 10毫秒
		appLogger.Warning("session.queueOut timeout")
	}
}

//in message
func (s *Session) dispathRaw(raw []byte) {
	var msg protocol.ClientComMessage
	appLogger.Debug("session.dispathRaw got '%s' from '%s'", raw, s.remoteAddr)
	if err := json.Unmarshal(raw, &msg); err != nil {
		//解析json出错
		s.queueOut(protocol.ErrMalformed("", "", time.Now().UTC().Round(time.Millisecond)))
		return
	}
	s.dispatch(&msg)
}

func (s *Session) dispatch(msg *protocol.ClientComMessage) {
	s.lastAction = time.Now().UTC().Round(time.Millisecond)

	msg.From = s.uid
	msg.Timestamp = s.lastAction

	switch {
	case msg.Hi != nil:
		s.hello(msg)
		break
	}
}

func (s *Session) hello(msg *protocol.ClientComMessage) {
	if msg.Hi.Version == "" {
		s.queueOut(protocol.ErrMalformed(msg.Hi.Id, "", msg.Timestamp))
		return
	}

	s.userAgent = msg.Hi.UserAgent
	s.deviceId = msg.Hi.DeviceID
	params := map[string]interface{}{
		"ver":   APP_VERSION,
		"build": buildstamp,
	}
	s.queueOut(&protocol.ServerComMessage{
		Ctrl: &protocol.MsgServerCtrl{
			Id:        msg.Hi.Id,
			Code:      http.StatusCreated,
			Text:      "created",
			Params:    params,
			Timestamp: msg.Timestamp,
		},
	})
}
