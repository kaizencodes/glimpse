package utils

import (
	"testing"
)

func TestFloatEquals(t *testing.T) {
	var tests = []struct {
		left     float64
		right    float64
		expected bool
	}{
		{1.000000001, 1.000000002, true},
		{1.0000001, 1.0000002, false},
	}

	for _, test := range tests {
		if result := FloatEquals(test.left, test.right); result != test.expected {
			t.Errorf("%f == %f result %t, expected %t", test.left, test.right, result, test.expected)
		}
	}
}

// Test that structs with the same values are considered equal.
func TestCompareStructs(t *testing.T) {
	type testStruct struct {
		Foo string
		Bar int
	}

	obj1 := testStruct{
		Foo: "foo",
		Bar: 1,
	}
	obj2 := obj1

	for _, diff := range Compare(obj1, obj2) {
		t.Errorf("Mismatch: %s", diff)
	}

	var tests = []struct {
		obj1 testStruct
		obj2 testStruct
	}{
		{testStruct{"foo", 1}, testStruct{"bar", 1}},
		{testStruct{"foo", 1}, testStruct{"foo", 2}},
		{testStruct{"", 2}, testStruct{"foo", 2}},
		{testStruct{}, testStruct{"foo", 2}},
	}
	for _, test := range tests {
		if diff := Compare(test.obj1, test.obj2); len(diff) == 0 {
			t.Errorf("Expected mismatch, but got none, for %v and %v", test.obj1, test.obj2)
		}
	}
}

func TestComparePointers(t *testing.T) {
	type testStruct struct {
		Foo string
		Bar int
	}

	obj1 := &testStruct{
		Foo: "foo",
		Bar: 1,
	}
	obj2 := obj1

	for _, diff := range Compare(obj1, obj2) {
		t.Errorf("Mismatch: %s", diff)
	}

	var tests = []struct {
		obj1 *testStruct
		obj2 *testStruct
	}{
		{&testStruct{"foo", 1}, &testStruct{"bar", 1}},
		{&testStruct{"foo", 1}, &testStruct{"foo", 2}},
		{&testStruct{"", 2}, &testStruct{"foo", 2}},
		{&testStruct{}, &testStruct{"foo", 2}},
	}
	for _, test := range tests {
		if diff := Compare(test.obj1, test.obj2); len(diff) == 0 {
			t.Errorf("Expected mismatch, but got none, for %v and %v", test.obj1, test.obj2)
		}
	}
}

func TestCompareDifferentSizeOfStructs(t *testing.T) {
	type testStruct1 struct {
		Foo string
		Bar int
	}
	type testStruct2 struct {
		Foo string
	}

	obj1 := testStruct1{
		Foo: "foo",
		Bar: 1,
	}
	obj2 := testStruct2{
		Foo: "foo",
	}

	if len(Compare(obj1, obj2)) == 0 {
		t.Errorf("Expected error for mismatching size of fields, but got none")
	}
}

// Test that struct with different fields are considered unequal.
func TestCompareDifferentStructs(t *testing.T) {
	type testStruct1 struct {
		Foo string
		Bar int
	}
	type testStruct2 struct {
		Foo string
		Baz int
	}

	obj1 := testStruct1{
		Foo: "foo",
		Bar: 1,
	}
	obj2 := testStruct2{
		Foo: "foo",
		Baz: 1,
	}

	if len(Compare(obj1, obj2)) == 0 {
		t.Errorf("Expected error for mismatching fields for Baz and Bar, but got none")
	}
}

// Test that slices with the same values are considered equal.
func TestSlices(t *testing.T) {
	obj1 := []float64{0.8, 0.5, 0.3}
	obj2 := obj1

	for _, diff := range Compare(obj1, obj2) {
		t.Errorf("Mismatch: %s", diff)
	}

	var tests = []struct {
		obj1 []float64
		obj2 []float64
	}{
		{[]float64{0.8, 0.5, 0.3}, []float64{0.8, 0.5, 0.4}},
		{[]float64{0.8, 0.5, 0.3}, []float64{0.8, 0.5}},
		{[]float64{0.8, 0.5}, []float64{}},
	}
	for _, test := range tests {
		if diff := Compare(test.obj1, test.obj2); len(diff) == 0 {
			t.Errorf("Expected mismatch, but got none, for %v and %v", test.obj1, test.obj2)
		}
	}
}

func TestCompareDifferentTypeOfSlices(t *testing.T) {
	obj1 := []float64{0.0, 1.0, 2.0}
	obj2 := []int{0, 1, 2}

	if len(Compare(obj1, obj2)) == 0 {
		t.Errorf("Expected error for mismatching type of slices, but got none")
	}
}

func TestCompareNestedStructs(t *testing.T) {
	type nestedStruct struct {
		Slice []int
	}

	type testStruct struct {
		Foo    string
		Nested nestedStruct
	}

	obj1 := testStruct{
		Foo: "foo",
		Nested: nestedStruct{
			Slice: []int{1, 2, 3},
		},
	}
	obj2 := obj1

	for _, diff := range Compare(obj1, obj2) {
		t.Errorf("Mismatch: %s", diff)
	}

	var tests = []struct {
		obj1 testStruct
		obj2 testStruct
	}{
		{
			testStruct{"foo", nestedStruct{[]int{1, 2, 3}}},
			testStruct{"foo", nestedStruct{[]int{1, 2, 4}}},
		},
		{
			testStruct{"foo", nestedStruct{[]int{1, 2, 3}}},
			testStruct{"bar", nestedStruct{[]int{1, 2, 3}}},
		},
		{
			testStruct{"foo", nestedStruct{[]int{1, 2, 3}}},
			testStruct{"foo", nestedStruct{[]int{1, 2}}},
		},
		{
			testStruct{"foo", nestedStruct{[]int{1, 2, 3}}},
			testStruct{"foo", nestedStruct{[]int{}}},
		},
		{
			testStruct{"foo", nestedStruct{[]int{1, 2, 3}}},
			testStruct{"foo", nestedStruct{}},
		},
		{
			testStruct{"foo", nestedStruct{[]int{1, 2, 3}}},
			testStruct{},
		},
	}
	for _, test := range tests {
		if diff := Compare(test.obj1, test.obj2); len(diff) == 0 {
			t.Errorf("Expected mismatch, but got none, for %v and %v", test.obj1, test.obj2)
		}
	}
}
