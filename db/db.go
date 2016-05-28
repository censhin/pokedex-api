package db

import (
        "log"
	"gopkg.in/mgo.v2"
)

var session *mgo.Session = nil
var host string = "172.17.0.2"
var port string = "27017"

func Connect() {
	var err error = nil
	s := host + ":" + port
        log.Println("Connecting to database " + s)
	session, err = mgo.Dial(s)
	if err != nil {
                log.Panic(err)
	}
}

func Session() *mgo.Session {
	return session.Copy()
}

func Close() {
	if session != nil {
		session.Close()
	}
}
