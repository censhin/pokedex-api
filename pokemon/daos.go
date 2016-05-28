package pokemon

import (
	"github.com/censhin/pokedex-api/db"
	"gopkg.in/mgo.v2/bson"
)

func CollectionDao() (Pokemons, error) {
	session := db.Session()
	defer session.Close()

	collection := session.DB("pokedex").C("pokemon")

	result := []Pokemon{}
	iter := collection.Find(nil).Iter()
	err := iter.All(&result)
	if err != nil {
		return Pokemons{}, err
	}

	return Pokemons{result}, nil
}

func MemberDao(name string) (Pokemon, error) {
	session := db.Session()
	defer session.Close()

	collection := session.DB("pokedex").C("pokemon")

	result := Pokemon{}
	err := collection.Find(bson.M{"name": name}).One(&result)
	if err != nil {
		return Pokemon{}, err
	}

	return result, nil
}

func CreateMemberDao(body Pokemon) error {
	session := db.Session()
	defer session.Close()

	collection := session.DB("pokedex").C("pokemon")

	return collection.Insert(body)
}

func UpdateMemberDao(name string, body map[string]interface{}) error {
	session := db.Session()
	defer session.Close()

	collection := session.DB("pokedex").C("pokemon")

	return collection.Update(bson.M{"name": name}, bson.M{"$set": body})
}

func DeleteMemberDao(name string) error {
	session := db.Session()
	defer session.Close()

	collection := session.DB("pokedex").C("pokemon")

	return collection.Remove(bson.M{"name": name})
}
