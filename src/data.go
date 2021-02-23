package randname

import (
	"math/rand"
	"strings"
)

type data map[string]map[string]float32

func createData(input *inputData, conf *dataConfig) *data {
	var newData *data
	if conf.append {
		newData = conf.currentData
	} else {
		newData = new(data)
	}

	var total int
	for c, v := range *input {
		total = len(v)
		for c2, v2 := range v {
			newData.calculate(c, c2, v2, total)
		}
	}

	return newData
}

func (d *data) calculate(k, c string, count, total int) {
	if *d == nil {
		*d = make(map[string]map[string]float32)
	}

	if _, ok := (*d)[k]; !ok {
		(*d)[k] = make(map[string]float32)
	}

	(*d)[k][c] = float32(count) / float32(total)
}

func (d *data) getPivot() string {
	index := rand.Intn(len(*d) - 1)

	count := 0
	for k := range *d {
		if count == index {
			return k
		}
		count++
	}

	return ""
}

func (d *data) getNext(pivot string) string {
	pivot = strings.ToLower(pivot)
	values := make([]*letterValue, 0)

	if _, ok := (*d)[pivot]; !ok {
		return ""
	}

	for c, v := range (*d)[pivot] {
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
