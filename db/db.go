package db

import (
	"log"
	"gopkg.in/mgo.v2"
	"github.com/censhin/pokedex-api/utils"
)

var session *mgo.Session = nil
var config = utils.GetConfig()
var host string = config.Db.Host
var port string = config.Db.Port

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
