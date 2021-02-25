package randname

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/pkg/errors"
)

var noChars *regexp.Regexp

type dataConfig struct {
	path        string
	append      bool
	currentData *data
}

func loadData(conf *dataConfig) (*data, error) {
	if !fileExists(conf.path) {
		return nil, fmt.Errorf("the source file %s doesn't exists", conf.path)
	}

	dataTemp := new(inputData)
	noChars = regexp.MustCompile(`(?mi)(?:[\p{L}\p{M}])`)

	dataChannel, err := readData(conf.path)
	if err != nil {
		return nil, errors.Wrapf(err, "trying lading data from file: %s", conf.path)
	}

	var letter, nextLetter, combine string
	for dataRead := range dataChannel {
		for i, c := range dataRead {
			letter = string(c)
			//skip no letters && check if it's no the last letter
			if !isValidChar(letter) || i+1 >= len(dataRead) {
				continue
			}

			//count match
			nextLetter = string(dataRead[i+1])
			dataTemp.increase(letter, nextLetter)

			if i+2 >= len(dataRead) {
				continue
			}

			combine = letter + nextLetter
			nextLetter = string(dataRead[i+2])
			dataTemp.increase(combine, nextLetter)
		}
	}

	return createData(dataTemp, conf), nil
}

//readData reads the data from the file, returns a channel where each line will be send
func readData(path string) (chan string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrapf(err, "trying open file: %s", path)
	}

	dataCh := make(chan string)
	go func(c chan string, file *os.File) {
		scanner := bufio.NewScanner(file)
		defer file.Close()
		defer close(c)

		for scanner.Scan() {
			c <- scanner.Text()
		}

		if err := scanner.Err(); err != nil {
			panic(err)
		}
	}(dataCh, file)

	return dataCh, nil
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
