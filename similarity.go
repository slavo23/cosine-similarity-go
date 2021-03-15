package similarity

import (
	"math"

	"github.com/slavo23/cosine-similarity-go/utils"
	"github.com/slavo23/cosine-similarity-go/vectormath"
)

func GetSimilarityCoefficent(str1 string, str2 string) float64 {
	v1, v2 := utils.Vectorize(str1, str2)

	v1Mag := vectormath.Magnitude(v1)
	v2Mag := vectormath.Magnitude(v2)

	similarityCoeff := vectormath.DotProduct(v1, v2) / (v1Mag * v2Mag)

	return 1.0 - 2.0*math.Acos(similarityCoeff)/math.Pi
}

func VectorizationTest(str1 string, str2 string) ([]float64, []float64) {
	v1, v2 := utils.Vectorize(str1, str2)

	return v1, v2
}
