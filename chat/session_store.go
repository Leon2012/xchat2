package main

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/Leon2012/xchat2/protocol"
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
	ss.rw.Lock()
	defer ss.rw.Unlock()
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
		s.sid = globals.idGen.GetStr()
	}
	ss.sessCache[s.sid] = &s
	return &s
}

func (ss *SessionStore) Get(sid string) *Session {
	ss.rw.Lock()
	defer ss.rw.Unlock()

	if s, ok := ss.sessCache[sid]; ok {
		return s
	}
	return nil
}

func (ss *SessionStore) Delete(s *Session) {
	ss.rw.Lock()
	defer ss.rw.Unlock()
	delete(ss.sessCache, s.sid)
}

func (ss *SessionStore) Shutdown() {
	ss.rw.Lock()
	defer ss.rw.Unlock()

	shutdown, _ := json.Marshal(protocol.NoErrShutdown(time.Now().UTC().Round(time.Millisecond)))
	for _, s := range ss.sessCache {
		if s.send != nil {
			s.send <- shutdown
		}
	}
}
