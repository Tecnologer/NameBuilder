package src

import (
	"math/rand"
	"strings"

	"github.com/pkg/errors"
)

func GetRandomName(l int) (string, error) {
	err := loadData("../src/input_names.txt")
	if err != nil {
		return "", errors.Wrapf(err, "creating random name with len of %d", l)
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
