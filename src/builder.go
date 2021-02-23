package randname

import (
	"path"
	"runtime"
	"strings"
	"unicode/utf8"

	"github.com/pkg/errors"
)

var currentData *data

func loadDefault() error {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("file 'input_names.txt' not found")
	}

	filepath := path.Join(path.Dir(filename), "../src/input_names.txt")

	data, err := loadData(newDataConfig(filepath, false))

	if err != nil {
		return err
	}

	currentData = data
	return nil
}

//AppendData appends data to the current data
func AppendData(path string) error {
	data, err := loadData(newDataConfig(path, true))
	if err != nil {
		return errors.Wrap(err, "append data fails")
	}

	currentData = data
	return nil
}

//LoadData replace the current data with the data in the specific path
func LoadData(path string) error {
	data, err := loadData(newDataConfig(path, false))
	if err != nil {
		return errors.Wrap(err, "append load new data fails")
	}

	currentData = data
	return nil
}

//GetRandomWithSeed generates a random name with the specified len and starting with the seed
func GetRandomWithSeed(seed string, l int) (string, error) {
	if currentData == nil {
		err := loadDefault()

		if err != nil {
			return "", errors.Wrap(err, "get random w/seed: data is not loaded")
		}
	}

	pivot := seed
	if seed == "" {
		for i := 0; i < 10 && pivot == ""; i++ {
			pivot = currentData.getPivot()
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
		pivot2 = currentData.getNext(combined)
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
	if currentData == nil {
		err := loadDefault()

		if err != nil {
			return "", errors.Wrap(err, "get random: data is not loaded")
		}
	}

	pivot := currentData.getPivot()
	return GetRandomWithSeed(pivot, l)
}

func newDataConfig(path string, append bool) *dataConfig {
	return &dataConfig{
		path:        path,
		append:      append,
		currentData: currentData,
	}
}
