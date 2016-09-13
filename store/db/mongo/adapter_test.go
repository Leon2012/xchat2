package mongo

import (
	"testing"

	"gopkg.in/mgo.v2/bson"

	"github.com/Leon2012/xchat2/store/types"
)

var config MongoConfig = MongoConfig{
	Url:     "127.0.0.1:27017",
	DbName:  "xchat",
	Timeout: 10,
}

var adapter *MongoAdapter

func getConn(t *testing.T) {
	adapter = NewMongoAdapter()
	err := adapter.Open(config)
	if err != nil {
		t.Error(err)
	}
}

func Test_Login(t *testing.T) {
	getConn(t)
	token := "xxxx"
	user, err := adapter.Login(token)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(user.Id)
	}
	adapter.Close()
}

func Test_Logout(t *testing.T) {
	getConn(t)
	user := types.User{}
	user.Token = "xxxx"
	err := adapter.Logout(&user)
	if err != nil {
		t.Error(err)
	}
	adapter.Close()
}

func Test_Struct(t *testing.T) {
	user := types.User{}
	t.Log(user)

	var user1 struct {
		Id    bson.ObjectId `bson:"_id,omitempty"`
		Uid   int64
		Uname string
		Token string
	}
	t.Log(user1)
}
