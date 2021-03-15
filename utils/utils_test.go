package utils

import (
	"reflect"
	"testing"
)

func TestWords(t *testing.T) {
	// Test for length
	s1 := "jack of all trades, master of none"

	expectedLength := 7
	gotLength := len(Words(s1))

	if gotLength != expectedLength {
		t.Errorf("Expected array of %d words, got %d", expectedLength, gotLength)
	}

	// Test contents of resulting array
	expectedArray := []string{"jack", "of", "all", "trades", "master", "of", "none"}
	gotArray := Words(s1)

	if !reflect.DeepEqual(gotArray, expectedArray) {
		t.Errorf("%q not the same as %s", gotArray, expectedArray)
	}

	// Test against empty string
	expectedArray = []string{""}
	gotArray = Words("")

	if !reflect.DeepEqual(gotArray, expectedArray) {
		t.Errorf("%q not the same as %s", gotArray, expectedArray)
	}

}

func TestVectorize(t *testing.T) {
	s1 := "test1 test2 test3"
	s2 := "test1 test2"

	v1, v2 := Vectorize(s1, s2)

	v1Len := len(v1)
	v2Len := len(v2)
	// Vector dimensions should always be equal, no matter the difference between strings
	if v1Len != v2Len {
		t.Errorf("v1 has length %d, but v2 has %d", v1Len, v2Len)
	}

	// Swapping sentences of different lengths does not result in dimensionality change
	v1, v2 = Vectorize(s2, s1)

	v1Len = len(v1)
	v2Len = len(v2)

	if v1Len != v2Len {
		t.Errorf("v1 has length %d, but v2 has %d", v1Len, v2Len)
	}

	expected1, expected2 := []float64{1, 1, 1}, []float64{1, 1, 0}

	if !reflect.DeepEqual(v1, expected1) {
		t.Errorf("%v does not equal %v,", v1, expected1)
	}

	if !reflect.DeepEqual(v2, expected2) {
		t.Errorf("%v does not equal %v", v2, expected2)
	}
}
