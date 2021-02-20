package randname

import (
	"bufio"
	"os"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

var data map[string]map[string]float32

func loadData(path string) error {
	dataTemp := make(map[string]map[string]int)
	noChars := regexp.MustCompile(`\W`)

	dataChannel, err := readData(path)
	if err != nil {
		return errors.Wrapf(err, "trying lading data from file: %s", path)
	}

	var letter, nextLetter, combine string
	for dataRead := range dataChannel {
		for i, c := range dataRead {
			letter = string(c)
			//skip no letters
			if noChars.Match([]byte(letter)) {
				continue
			}

			letter = strings.ToLower(letter)

			//check if it's no the last letter
			if i+1 >= len(dataRead) {
				continue
			}

			//initialize the map for the letter
			if _, ok := dataTemp[letter]; !ok {
				dataTemp[letter] = make(map[string]int)
			}

			//count match
			nextLetter = strings.ToLower(string(dataRead[i+1]))
			dataTemp[letter][nextLetter]++

			if i+2 >= len(dataRead) {
				continue
			}

			combine = letter + nextLetter
			if _, ok := dataTemp[combine]; !ok {
				dataTemp[combine] = make(map[string]int)
			}

			nextLetter = strings.ToLower(string(dataRead[i+2]))
			dataTemp[combine][nextLetter]++
		}
	}

	calculatePercentage(dataTemp)

	return nil
}

func readData(path string) (chan []byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrapf(err, "trying open file: %s", path)
	}

	dataCh := make(chan []byte)
	go func(c chan []byte, file *os.File) {
		scanner := bufio.NewScanner(file)
		defer file.Close()
		defer close(c)

		for scanner.Scan() {
			c <- scanner.Bytes()
		}

		if err := scanner.Err(); err != nil {
			panic(err)
		}
	}(dataCh, file)

	return dataCh, nil
}

func calculatePercentage(input map[string]map[string]int) {
	data = make(map[string]map[string]float32)
	var total float32
	for c, v := range input {
		if _, ok := data[c]; !ok {
			data[c] = make(map[string]float32)
		}

		total = float32(len(v))
		for c2, v2 := range v {
			data[c][c2] = float32(v2) / total
		}
	}
}
