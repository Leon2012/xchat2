package mongo

import (
	"errors"
	"time"

	"github.com/Leon2012/xchat2/store"
	"github.com/Leon2012/xchat2/store/types"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func init() {
	store.Register("mongo", NewMongoAdapter())
}

type MongoAdapter struct {
	Conf    MongoConfig
	Session *mgo.Session
	Db      *mgo.Database
	isOpen  bool
}

type MongoConfig struct {
	Url      string
	DbName   string
	Timeout  int
	Username string
	Password string
}

func NewMongoAdapter() *MongoAdapter {
	return &MongoAdapter{}
}

func (a *MongoAdapter) Open(c interface{}) error {
	config, ok := c.(MongoConfig)
	if !ok {
		return errors.New("Invaid Config")
	}
	timeout := time.Duration(config.Timeout) * time.Second
	session, err := mgo.DialWithTimeout(config.Url, timeout)
	if err != nil {
		return err
	}
	session.SetMode(mgo.Monotonic, true)
	db := session.DB(config.DbName)
	if config.Username != "" {
		err := db.Login(config.Username, config.Password)
		if err != nil {
			return err
		}
	}
	a.Db = db
	a.Session = session
	a.Conf = config
	a.isOpen = true
	return nil
}

func (a *MongoAdapter) Close() error {
	a.Session.Close()
	a.isOpen = false
	return nil
}

func (a *MongoAdapter) IsOpen() bool {
	return a.isOpen
}

func (a *MongoAdapter) Login(token string) (*types.User, error) {
	user := types.User{}
	users := a.Db.C("users")
	err := users.Find(bson.M{"token": token}).One(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (a *MongoAdapter) Logout(user *types.User) error {
	users := a.Db.C("users")
	return users.Remove(bson.M{"token": user.Token})
}
