package randname

import (
	"fmt"
	"math/rand"
	"path"
	"runtime"
	"strings"
	"unicode/utf8"
)

func init() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("file 'input_names.txt' not found")
	}

	filepath := path.Join(path.Dir(filename), "../src/input_names.txt")

	err := loadData(filepath)
	if err != nil {
		panic(err)
	}

}

//GetRandomWithSeed generates a random name with the specified len and starting with the seed
func GetRandomWithSeed(seed string, l int) (string, error) {
	if data == nil {
		return "", fmt.Errorf("data is not loaded")
	}

	pivot := seed
	if seed == "" {
		for i := 0; i < 10 && pivot == ""; i++ {
			pivot = getPivot()
		}
	}

	var pivot2 string
	name := strings.ToUpper(pivot[:1])
	if utf8.RuneCount([]byte(pivot)) > 1 {
		name += pivot[1:]
	}

	if utf8.RuneCount([]byte(pivot)) > 2 {
		pivot = pivot[utf8.RuneCount([]byte(pivot))-2:]
	}
	combined := pivot

	for i := utf8.RuneCount([]byte(seed)); i < l; i++ {
		pivot2 = getNext(combined)
		if pivot == "" {
			continue
		}

		if utf8.RuneCount([]byte(pivot)) > 1 {
			combined = pivot[1:] + pivot2
		} else {
			combined = pivot + pivot2
		}

		pivot = pivot2
		name += string(pivot)
	}

	return name, nil
}

//GetRandom generates a random name with the specified len
func GetRandom(l int) (string, error) {
	if data == nil {
		return "", fmt.Errorf("data is not loaded")
	}

	pivot := getPivot()
	return GetRandomWithSeed(pivot, l)
}

func getPivot() string {
	index := rand.Intn(len(data) - 1)

	count := 0
	for k := range data {
		if count == index {
			return k
		}
		count++
	}

	return ""
}

func getNext(pivot string) string {
	pivot = strings.ToLower(pivot)
	values := make([]*letterValue, 0)

	if _, ok := data[pivot]; !ok {
		return ""
	}

	for c, v := range data[pivot] {
		values = append(values, &letterValue{c, v})
	}

	By(byProbilityDesc).Sort(values)

	if len(values) > 5 {
		index := rand.Intn(4)
		return values[index].letter
	}

	index := 0
	if len(values) > 1 {
		index = rand.Intn(len(values) - 1)
	}

	return values[index].letter

}
