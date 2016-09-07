package main

import "sync"

type sessionStore struct {
	rw            sync.RWMutex
	sessions      map[int64]*session //uid:session
	serverCounter map[int32]int32    // server->count
	userCounter   map[int64]int32
}

func newSessionStore() *sessionStore {
	return &sessionStore{
		sessions:      make(map[int64]*session),
		serverCounter: make(map[int32]int32),
		userCounter:   make(map[int64]int32),
	}
}

func (s *sessionStore) Put(uid int64, sid int32) {
	var (
		session *session
		ok      bool
	)
	s.rw.Lock()
	defer s.rw.Unlock()
	if session, ok = s.sessions[uid]; !ok {
		s.sessions[uid] = newSession()
	}
}
