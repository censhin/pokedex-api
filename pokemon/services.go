package pokemon

func CollectionService() Pokemons {
	return CollectionDao()
}

func MemberService(name string) Pokemon {
	return MemberDao(name)
}
