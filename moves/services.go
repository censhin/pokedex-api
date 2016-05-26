package moves

func CollectionService() (Moves) {
	return CollectionDao()
}

func MemberService(name string) (Move) {
	return MemberDao(name)
}
