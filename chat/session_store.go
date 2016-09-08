package main

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type SessionStore struct {
	rw        sync.RWMutex
	sessCache map[string]*Session
	lifeTime  time.Duration
}

func newSessionStore(lifetime time.Duration) *SessionStore {
	store := &SessionStore{
		lifeTime:  lifetime,
		sessCache: make(map[string]*Session),
	}
	return store
}

func (ss *SessionStore) Create(conn interface{}, sid string) *Session {
	var s Session
	s.sid = sid
	switch c := conn.(type) {
	case *websocket.Conn:
		s.proto = WEBSOCK
		s.ws = c
		break
	default:
		s.proto = NONE
	}
	if s.proto != NONE {
		s.send = make(chan []byte, 64)
		s.stop = make(chan []byte, 1)
		s.detach = make(chan string, 64)
	}
	s.lastTouched = time.Now()
	if s.sid == "" {
		s.sid = ""
	}

	ss.rw.Lock()
	ss.sessCache[s.sid] = &s
	ss.rw.Unlock()

	return &s
}
