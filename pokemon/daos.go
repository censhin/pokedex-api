package pokemon

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

func PokemonsDao() Pokemons {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	collection := session.DB("pokedex").C("pokemon")

	result := []Pokemon{}
	iter := collection.Find(nil).Iter()
	err = iter.All(&result)
	if err != nil {
		log.Fatal(err)
	}

	return Pokemons{result}
}

func PokemonDao(name string) Pokemon {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	collection := session.DB("pokedex").C("pokemon")

	result := Pokemon{}
	err = collection.Find(bson.M{"name": name}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}
