package vectormath

import (
	"math"
	"testing"
)

var v1 = []float64{1, 0, 2, 5}
var v2 = []float64{0, 3, 1, 4}

func lessOrMoreThan(a float64, pred float64) bool {
	return a < pred || a > pred
}

func TestDotProduct(t *testing.T) {
	dot := DotProduct(v1, v2)
	expectedDot := 22.0

	if dot != expectedDot {
		t.Errorf("Expected %3.6f as dot product, got %3.6f", expectedDot, dot)
	}
}

func TestMagnitude(t *testing.T) {
	mag1 := Magnitude(v1)
	mag2 := Magnitude(v2)

	expected1 := math.Sqrt(30)
	expected2 := math.Sqrt(26)

	if lessOrMoreThan(mag1, expected1) {
		t.Errorf("Expected %3.6f as vector length, got %3.6f", expected1, mag1)
	}

	if lessOrMoreThan(mag2, expected2) {
		t.Errorf("Expected %3.6f as vector length, got %3.6f", expected1, mag1)
	}
}
