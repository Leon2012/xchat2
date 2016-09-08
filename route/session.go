package main

type session struct {
	uid int64 //User Id
	sid int32 //Server Id
	seq int32 //seq
}

func newSession() *session {
	return &session{}
}

func (s *session) Put(uid int64, sid int32) {
	s.uid = uid
	s.sid = sid
}
