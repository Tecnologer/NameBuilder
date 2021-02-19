package src

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

func init() {
	err := loadData("../../src/input_names.txt")
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
	values := make([]*letterValue, 0)

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

type letterValue struct {
	letter      byte
	probability float32
}

type letterSorter struct {
	letters []*letterValue
	by      By
}

func byProbilityDesc(l, l2 *letterValue) bool {
	return l.probability > l2.probability
}

// By is the type of a "less" function that defines the ordering of its Planet arguments.
type By func(p1, p2 *letterValue) bool

// Sort is a method on the function type, By, that sorts the argument slice according to the function.
func (by By) Sort(letters []*letterValue) {
	ps := &letterSorter{
		letters: letters,
		by:      by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(ps)
}

// Len is part of sort.Interface.
func (s *letterSorter) Len() int {
	return len(s.letters)
}

// Swap is part of sort.Interface.
func (s *letterSorter) Swap(i, j int) {
	s.letters[i], s.letters[j] = s.letters[j], s.letters[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *letterSorter) Less(i, j int) bool {
	return s.by(s.letters[i], s.letters[j])
}
