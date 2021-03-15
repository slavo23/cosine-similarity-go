package utils

import (
	"log"
	"regexp"
	"strings"

	"github.com/iancoleman/orderedmap"
)

func Words(str string) []string {

	reg, err := regexp.Compile("[^a-zA-Z0-9]+")

	if err != nil {
		log.Fatal(err)
	}

	replaced := reg.ReplaceAllString(str, " ")

	return strings.Split(replaced, " ")
}

func MapValuesAsVector(orderedMap orderedmap.OrderedMap) []float64 {
	vec := []float64{}

	for _, key := range orderedMap.Keys() {
		scalar, _ := orderedMap.Get(key)

		vec = append(vec, scalar.(float64))
	}

	return vec
}

func Vectorize(str1 string, str2 string) ([]float64, []float64) {
	table1 := orderedmap.New()
	table2 := orderedmap.New()

	var firstString string
	var secondString string

	if len(str2) > len(str1) {
		table1, table2 = table2, table1
		firstString, secondString = str2, str1
	} else {
		firstString, secondString = str1, str2
	}

	for _, word := range Words(firstString) {
		currentCount, ok := table1.Get(word)

		table2.Set(word, float64(0))

		if !ok {
			table1.Set(word, float64(1))
		} else {
			table1.Set(word, currentCount.(float64)+1)
		}
	}

	for _, word := range Words(secondString) {
		_, ok := table1.Get(word)
		currentCount, _ := table2.Get(word)

		if ok {
			table2.Set(word, currentCount.(float64)+1)
		}
	}

	return MapValuesAsVector(*table1), MapValuesAsVector(*table2)
}
