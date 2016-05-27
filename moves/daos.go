package moves

import (
	"log"
	"gopkg.in/mgo.v2/bson"
	"github.com/censhin/pokedex-api/db"
)

func CollectionDao() Moves {
	session := db.Session()
	defer session.Close()

	collection := session.DB("pokedex").C("moves")

	result := []Move{}
	iter := collection.Find(nil).Iter()
	err := iter.All(&result)
	if err != nil {
		log.Fatal(err)
	}

	return Moves{result}
}

func MemberDao(name string) Move {
	session := db.Session()
	defer session.Close()

	collection := session.DB("pokedex").C("moves")

	result := Move{}
	err := collection.Find(bson.M{"name": name}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}
