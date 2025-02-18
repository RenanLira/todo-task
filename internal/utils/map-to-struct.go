package utils

import (
	"reflect"
)

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return string(s[0]-32) + s[1:]
}

func MapToStruct(data map[string]any, result any) error {
	val := reflect.ValueOf(result).Elem()

	for key, value := range data {
		field := val.FieldByNameFunc(func(s string) bool {
			// Ignora case ao comparar nomes
			return s == key || s == capitalize(key)
		})
		if field.IsValid() && field.CanSet() {
			valValue := reflect.ValueOf(value)
			if valValue.IsValid() {
				field.Set(valValue)
			}
		}
	}
	return nil
}
