package pokemon


func CollectionService() (Pokemons, error) {
	return CollectionDao()
}

func PostCollectionService(body Pokemon) error {
	return CreateMemberDao(body)
}

func MemberService(name string) (Pokemon, error) {
	return MemberDao(name)
}

func PutMemberService(name string, body map[string]interface{}) error {
	return UpdateMemberDao(name, body)
}
