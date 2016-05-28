package pokemon

import (
	"log"
	"github.com/censhin/pokedex-api/db"
	"gopkg.in/mgo.v2/bson"
)

func CollectionDao() Pokemons {
	session := db.Session()
	defer session.Close()

	collection := session.DB("pokedex").C("pokemon")

	result := []Pokemon{}
	iter := collection.Find(nil).Iter()
	err := iter.All(&result)
	if err != nil {
		log.Fatal(err)
	}

	return Pokemons{result}
}

func MemberDao(name string) Pokemon {
	session := db.Session()
	defer session.Close()

	collection := session.DB("pokedex").C("pokemon")

	result := Pokemon{}
	err := collection.Find(bson.M{"name": name}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func UpdateMemberDao(name string, body Pokemon) error {
	session := db.Session()
	defer session.Close()

	collection := session.DB("pokedex").C("pokemon")

	return collection.Update(name, body)
}
