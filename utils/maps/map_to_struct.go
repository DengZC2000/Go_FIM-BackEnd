package maps

import (
	"reflect"
)

func MapToStruct(data map[string]any, dst any) {
	//dst必须是一个指针
	t := reflect.TypeOf(dst).Elem()
	v := reflect.ValueOf(dst).Elem()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("json")
		if tag == "" || tag == "-" {
			continue
		}
		mapField, ok := data[tag]
		if !ok {
			continue
		}
		val := v.Field(i)
		if field.Type.Kind() == reflect.Ptr {
			switch field.Type.Elem().Kind() {
			case reflect.String:
				mapFieldValue := reflect.ValueOf(mapField)
				if mapFieldValue.Type().Kind() == reflect.String {
					strVal := mapField.(string)
					val.Set(reflect.ValueOf(&strVal))
				}
			}
		}
	}
}
