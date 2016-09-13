package types

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type ObjHeader struct {
	Id        bson.ObjectId `bson:"_id,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type User struct {
	ObjHeader `bson:",inline"` //此处必需用inline, 见https://godoc.org/gopkg.in/mgo.v2/bson
	Uid       int64
	Uname     string
	Token     string
}

type Topic struct {
	ObjHeader `bson:",inline"`
	Name      string
	SeqId     int64
	Owner     int64
}

type Subscription struct {
	ObjHeader `bson:",inline"`
	Topic     string
	User      int64
	ReadSeqId int64
	RecvSeqId int64
}

type Message struct {
	ObjHeader `bson:",inline"`
	SeqId     int
	Topic     string
	From      string
	Content   interface{}
}
