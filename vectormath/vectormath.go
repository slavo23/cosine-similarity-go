package vectormath

import (
	"math"
)

func DotProduct(vec1 []float64, vec2 []float64) float64 {

	var product float64 = 0.0

	for i := range vec1 {
		product += vec1[i] * vec2[i]
	}

	return product
}

func Magnitude(vec []float64) float64 {
	return math.Sqrt(DotProduct(vec, vec))
}
