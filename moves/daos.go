package moves

import (
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func MovesDao() Moves {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	collection := session.DB("pokedex").C("moves")

	result := []Move{}
	iter := collection.Find(nil).Iter()
	err = iter.All(&result)
	if err != nil {
		log.Fatal(err)
	}

	return Moves{result}
}

func MoveDao(name string) Move {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	collection := session.DB("pokedex").C("moves")

	result := Move{}
	err = collection.Find(bson.M{"name": name}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}
