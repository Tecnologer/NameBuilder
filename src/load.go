package src

import (
	"bufio"
	"os"
	"regexp"

	"github.com/pkg/errors"
)

var data map[byte]map[byte]float32

func loadData(path string) error {
	dataTemp := make(map[byte]map[byte]int)
	noChars := regexp.MustCompile(`\W`)

	dataChannel, err := readData(path)
	if err != nil {
		return errors.Wrapf(err, "trying lading data from file: %s", path)
	}

	for dataRead := range dataChannel {
		for i, c := range dataRead {
			if noChars.Match([]byte{c}) {
				continue
			}

			if _, ok := dataTemp[c]; !ok {
				dataTemp[c] = make(map[byte]int)
			}

			if i+1 < len(dataRead) {
				dataTemp[c][dataRead[i+1]]++
			}
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

func calculatePercentage(input map[byte]map[byte]int) {
	data = make(map[byte]map[byte]float32)
	var total float32
	for c, v := range input {
		if _, ok := data[c]; !ok {
			data[c] = make(map[byte]float32)
		}

		total = float32(len(v))
		for c2, v2 := range v {
			data[c][c2] = float32(v2) / total
		}
	}
}
