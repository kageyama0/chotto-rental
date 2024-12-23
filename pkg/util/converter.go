package util

import (
	"reflect"
	"strings"
)

// StructToMap converts a struct to a map using reflection.
// It uses json tags as keys if available, otherwise uses field names.
func StructToMap(obj interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	val := reflect.ValueOf(obj)
	typ := val.Type()

	if typ.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = val.Type()
	}

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag.Get("json")
		if tag == "" {
			tag = field.Name
		}
		if tag == "-" {
			continue
		}
		if idx := strings.Index(tag, ","); idx != -1 {
			tag = tag[:idx]
		}
		result[tag] = val.Field(i).Interface()
	}

	return result
}
