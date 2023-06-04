// utility functions that are used throughout the project
package utils

import (
	"fmt"
	"reflect"
)

const EPSILON float64 = 0.00000001

// FloatEquals compares two floats, used to avoid floating point errors.
// Round-off error can make two numbers that should be equivalent instead be slightly different.
func FloatEquals(a, b float64) bool {
	if (a-b) < EPSILON && (b-a) < EPSILON {
		return true
	}
	return false
}

// Compare compares two objects and returns a list of differences. Used for testing.
func Compare(obj1, obj2 interface{}) []string {
	diff := []string{}

	v1 := reflect.ValueOf(obj1)
	v2 := reflect.ValueOf(obj2)

	if v1.Kind() != v2.Kind() {
		return []string{"Incompatible types: " + v1.Kind().String() + " " + v2.Kind().String()}
	}

	switch v1.Kind() {
	case reflect.Struct:
		if v1.NumField() != v2.NumField() {
			diff = append(diff, fmt.Sprintf("Field number mismatch: actual %d != expected %d", v1.NumField(), v2.NumField()))
			return diff
		}

		t1 := reflect.TypeOf(obj1)
		t2 := reflect.TypeOf(obj2)

		for i := 0; i < v1.NumField(); i++ {
			fieldName1 := t1.Field(i).Name
			fieldName2 := t2.Field(i).Name

			if fieldName1 != fieldName2 {
				diff = append(diff, fmt.Sprintf("Field name mismatch: actual %s != expected %s", t1.Field(i).Name, t2.Field(i).Name))
				return diff
			}

			field1 := v1.Field(i)
			field2 := v2.Field(i)

			// If the field is a struct, recurse then append the field name to the diff
			if complexObject(field1) {
				nestedDiff := Compare(field1.Interface(), field2.Interface())
				for _, diffStr := range nestedDiff {
					// using the first fields name, since they should be the same
					diff = append(diff, fmt.Sprintf("%s.%s", fieldName1, diffStr))
				}
			} else if !reflect.DeepEqual(field1.Interface(), field2.Interface()) {
				diff = append(diff, fmt.Sprintf("%s: actual %v != expected %v", fieldName1, field1.Interface(), field2.Interface()))
			}
		}
	case reflect.Slice, reflect.Array:
		if v1.Len() != v2.Len() {
			diff = append(diff, fmt.Sprintf("Length mismatch: actual %d != expected %d", v1.Len(), v2.Len()))
			return diff
		}

		for i := 0; i < v1.Len(); i++ {
			field1 := v1.Index(i)
			field2 := v2.Index(i)

			// If the field is a struct, recurse then append the field name to the diff
			if complexObject(field1) {
				nestedDiff := Compare(field1.Interface(), field2.Interface())
				for _, diffStr := range nestedDiff {
					diff = append(diff, fmt.Sprintf("%d.%s", i, diffStr))
				}
			} else if !reflect.DeepEqual(field1.Interface(), field2.Interface()) {
				// using the first fields name, since they should be the same
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
