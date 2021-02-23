package randname

import "strings"

type inputData map[string]map[string]int

func (d *inputData) increase(k, c string) {

	k = strings.ToLower(k)
	c = strings.ToLower(c)

	if !isValidChar(c) {
		return
	}

	if *d == nil {
		*d = make(map[string]map[string]int)
	}

	//initialize the map for the letter
	if _, ok := (*d)[k]; !ok {
		(*d)[k] = make(map[string]int)
	}

	(*d)[k][c]++
}

func isValidChar(c string) bool {
	return noChars.Match([]byte(c))
}
