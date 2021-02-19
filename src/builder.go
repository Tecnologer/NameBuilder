package src

import (
	"fmt"
	"math/rand"
	"strings"
)

func init() {
	err := loadData("../src/input_names.txt")
	if err != nil {
		panic(err)
	}
}
func GetRandomName(l int) (string, error) {
	if data == nil {
		return "", fmt.Errorf("data is not loaded")
	}

	pivot := getFirstLetter()
	name := strings.ToUpper(string(pivot))
	for i := 1; i < l; i++ {
		pivot = getNext(pivot)
		name += strings.ToLower(string(pivot))
	}
	return name, nil
}

func getFirstLetter() byte {
	index := rand.Intn(len(data) - 1)

	count := 0
	for k := range data {
		if count == index {
			return k
		}
		count++
	}

	return 0
}

func getNext(pivot byte) byte {
	var value float32
	var letter byte

	for c, v := range data[pivot] {
		if v > value {
			value = v
			letter = c
		}
	}

	return letter
}
