package models

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

type (
	M bson.M
)

var (
	notfound = "not found"
)

func ConnectMgo(mgourl string) (*mgo.Session, error) {
	session, err := mgo.DialWithTimeout(mgourl, 2*time.Second)
	if nil == err {
		session.SetSyncTimeout(2 * time.Second)
		session.SetSocketTimeout(2 * time.Second)
	}
	if nil != err {
		return session, err
	}
	session.SetMode(mgo.Monotonic, true)
	return session, err
}
