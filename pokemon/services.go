package pokemon

func PokemonsService() Pokemons {
	return PokemonsDao()
}

func PokemonService(name string) Pokemon {
	return PokemonDao(name)
}
