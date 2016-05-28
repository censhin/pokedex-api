package pokemon

import (
	"gopkg.in/mgo.v2"
)

func CollectionService() CollectionResponse {
	res, err := CollectionDao()
	if err == mgo.ErrNotFound {
		return CollectionResponse{Code: 404, Msg: err.Error()}
	} else if err != nil {
		return CollectionResponse{Code: 400, Msg: err.Error()}
	} else {
		return CollectionResponse{Code: 200, Msg: "OK", Body: res}
	}
}

func PostCollectionService(body Pokemon) CollectionResponse {
	err := CreateMemberDao(body)
	if err != nil {
		return CollectionResponse{Code: 400, Msg: err.Error()}
	} else {
		return CollectionResponse{Code: 204, Msg: "OK", Body: Pokemons{}}
	}
}

func MemberService(name string) Response {
	res, err := MemberDao(name)
	if err == mgo.ErrNotFound {
		return Response{Code: 404, Msg: err.Error()}
	} else if err != nil {
		return Response{Code: 400, Msg: err.Error()}
	} else {
		return Response{Code: 200, Msg: "OK", Body: res}
	}
}

func PutMemberService(name string, body map[string]interface{}) Response {
	err := UpdateMemberDao(name, body)
	if err != nil {
		return Response{Code: 400, Msg: err.Error()}
	} else {
		return Response{Code: 204, Msg: "OK", Body: Pokemon{}}
	}
}

func DeleteMemberService(name string) Response {
	err := DeleteMemberDao(name)
	if err != nil {
		return Response{Code: 400, Msg: err.Error()}
	} else {
		return Response{Code: 204, Msg: "OK", Body: Pokemon{}}
	}
}
