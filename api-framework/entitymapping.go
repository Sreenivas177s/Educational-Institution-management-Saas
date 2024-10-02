package apiframework

import (
	"chat-server/api-framework/entity"
	"reflect"
)

var entityMapping map[string]reflect.Type = map[string]reflect.Type{
	"chats": reflect.TypeOf(entity.Chat{}),
}

func GetDefinedType(dtype string) reflect.Type {
	if dtype == "" {
		return nil
	}
	return entityMapping[dtype]
}
