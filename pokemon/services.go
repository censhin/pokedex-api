package pokemon

import (
	"log"
	"fmt"
)

func CollectionService() Pokemons {
	return CollectionDao()
}

func MemberService(name string) Pokemon {
	return MemberDao(name)
}

func PutMemberService(name string, body map[string]interface{}) error {
	//err := UpdateMemberDao(name, body)
	fmt.Println(body)
	var err error = nil
	if err != nil {
		log.Fatal(err)
		return err
	} else {
		return nil
	}
}
