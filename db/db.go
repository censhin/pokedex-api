package db

import (
	"gopkg.in/mgo.v2"
)

var session *mgo.Session = nil
var host string = "localhost"
var port string = "27017"

func Connect() {
	var err error = nil
	s := host + ":" + port
	session, err = mgo.Dial(s)
	if err != nil {
		panic(err)
	}
}

func Session() *mgo.Session {
	return session
}

func Close() {
	if session != nil {
		session.Close()
	}
}