package utils

import (
	"fmt"
	"reflect"
)

// Compare compares two objects and returns a list of differences
func Compare(obj1, obj2 interface{}) []string {
	diff := []string{}

	v1 := reflect.ValueOf(obj1)
	v2 := reflect.ValueOf(obj2)

	if v1.Kind() != v2.Kind() {
		return []string{"Incompatible types: " + v1.Kind().String() + " " + v2.Kind().String()}
	}

	switch v1.Kind() {
	case reflect.Struct:
		for i := 0; i < v1.NumField(); i++ {
			field1 := v1.Field(i)
			field2 := v2.Field(i)

			if complexObject(field1) {
				nestedDiff := Compare(field1.Interface(), field2.Interface())
				for _, diffStr := range nestedDiff {
					fieldName := v1.Type().Field(i).Name
					diff = append(diff, fmt.Sprintf("%s.%s", fieldName, diffStr))
				}
			} else if !reflect.DeepEqual(field1.Interface(), field2.Interface()) {
				fieldName := v1.Type().Field(i).Name
				diff = append(diff, fmt.Sprintf("%s: actual %v != expected %v", fieldName, field1.Interface(), field2.Interface()))
			}
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < v1.Len(); i++ {
			field1 := v1.Index(i)
			field2 := v2.Index(i)

			if complexObject(field1) {
				nestedDiff := Compare(field1.Interface(), field2.Interface())
				for _, diffStr := range nestedDiff {
					diff = append(diff, fmt.Sprintf("%d.%s", i, diffStr))
				}
			} else if !reflect.DeepEqual(field1.Interface(), field2.Interface()) {
				diff = append(diff, fmt.Sprintf("%d: actual %v != expected %v", i, field1.Interface(), field2.Interface()))
			}
		}
	}

	return diff
}

func complexObject(object reflect.Value) bool {
	complex := false

	switch object.Kind() {
	case reflect.Struct:
		complex = true
	case reflect.Slice, reflect.Array:
		if object.Len() > 0 && object.Index(0).Kind() == reflect.Struct {
			complex = true
		}
	}
	return complex
}
