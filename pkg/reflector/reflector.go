package reflector

import (
	"reflect"
	"strconv"
)

type FieldValue struct {
	FieldName *string
	Value     string
}

type ValuesMap map[string][]FieldValue

func GetFieldsValuesAndNames(entity any) (names []string, values ValuesMap) {
	values = make(ValuesMap)

	typeof := reflect.TypeOf(entity)
	valueof := reflect.ValueOf(entity)
	for i := 0; i < typeof.NumField(); i++ {
		name := typeof.Field(i).Name
		names = append(names, name)

		value := valueof.Field(i)
		values[name] = append(values[name], GetValueMap(typeof.Field(i), value)...)
	}

	return
}

func GetValueMap(typeof reflect.StructField, valueof reflect.Value) (values []FieldValue) {
	switch valueof.Kind() {
	case reflect.String:
		fieldValue := FieldValue{Value: valueof.String()}
		values = append(values, fieldValue)
	case reflect.Int:
		fieldValue := FieldValue{Value: strconv.FormatInt(valueof.Int(), 10)}
		values = append(values, fieldValue)
	case reflect.Struct:
		v := reflect.ValueOf(valueof.Interface())
		t := typeof.Type
		for i := 0; i < t.NumField(); i++ {
			fieldName := t.Field(i).Name
			fieldValue := FieldValue{FieldName: &fieldName, Value: v.Field(i).String()}
			values = append(values, fieldValue)
		}
	}

	return
}

func GetField(entity any) (names []string) {
	r := reflect.TypeOf(entity)
	for i := 0; i < r.NumField(); i++ {
		f := r.Field(i)
		names = append(names, f.Name)
	}

	return
}
