package main

type session struct {
	uid int64 //User Id
	sid int32 //Server Id
	seq int32 //seq
}

func newSession() *session {
	return &session{}
}
