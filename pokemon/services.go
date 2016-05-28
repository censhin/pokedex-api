package pokemon

import (
	"gopkg.in/mgo.v2"
)

func CollectionService() Response {
	res, err := CollectionDao()
	if err == mgo.ErrNotFound {
		return Response{Code: 404, Body: err.Error()}
	} else if err != nil {
		return Response{Code: 400, Body: err.Error()}
	} else {
		return Response{Code: 200, Body: res}
	}
}

func PostCollectionService(body Pokemon) Response {
	err := CreateMemberDao(body)
	if err != nil {
		return Response{Code: 400, Body: err.Error()}
	} else {
		return Response{Code: 204, Body: nil}
	}
}

func MemberService(name string) Response {
	res, err := MemberDao(name)
	if err == mgo.ErrNotFound {
		return Response{Code: 404, Body: err.Error()}
	} else if err != nil {
		return Response{Code: 400, Body: err.Error()}
	} else {
		return Response{Code: 200, Body: res}
	}
}

func PutMemberService(name string, body map[string]interface{}) Response {
	err := UpdateMemberDao(name, body)
	if err != nil {
		return Response{Code: 400, Body: err.Error()}
	} else {
		return Response{Code: 204, Body: nil}
	}
}

func DeleteMemberService(name string) Response {
	err := DeleteMemberDao(name)
	if err != nil {
		return Response{Code: 400, Body: err.Error()}
	} else {
		return Response{Code: 204, Body: nil}
	}
}
