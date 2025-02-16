package utils

import (
	"reflect"
)

func UpdateIfNotNull[T any](target *T, value *T) {
	if value != nil {
		*target = *value
	}
}

func CopyStruct[T interface{}](target *T, value interface{}) {
	valueValue := reflect.ValueOf(value)
	valueType := reflect.TypeOf(value)

	for i := 0; i < valueValue.NumField(); i++ {
		field := valueType.Field(i)
		fieldValue := valueValue.Field(i)

		if fieldValue.IsNil() {
			continue
		}

		reflect.ValueOf(target).Elem().FieldByName(field.Name).Set(fieldValue.Elem())

	}

}
