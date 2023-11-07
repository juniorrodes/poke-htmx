package reflector

import (
	"reflect"
)

func GetFieldsNames(entity any) (names []string) {
	typeof := reflect.TypeOf(entity)
	for i := 0; i < typeof.NumField(); i++ {
		if typeof.Field(i).Tag.Get("reflect") == "ignore" {
			continue
		}
		name := typeof.Field(i).Name
		names = append(names, name)
	}

	return
}
