package pokemon

import (
	"log"
	"github.com/censhin/pokedex-api/db"
	"gopkg.in/mgo.v2/bson"
)

func PokemonsDao() Pokemons {
	session := db.Session()

	collection := session.DB("pokedex").C("pokemon")

	result := []Pokemon{}
	iter := collection.Find(nil).Iter()
	err := iter.All(&result)
	if err != nil {
		log.Fatal(err)
	}

	return Pokemons{result}
}

func PokemonDao(name string) Pokemon {
	session := db.Session()

	collection := session.DB("pokedex").C("pokemon")

	result := Pokemon{}
	err := collection.Find(bson.M{"name": name}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}