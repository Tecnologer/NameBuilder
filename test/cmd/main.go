package main

import (
	"flag"
	"fmt"
	"math/rand"

	"github.com/tecnologer/NameBuilder/src"
)

var (
	count      = flag.Int("count", 1, "how many names should be generated")
	nameLen    = flag.Int("len", 4, "length of the name(s)")
	minNameLen = flag.Int("len-min", 4, "the minimum of lenght of the name(s)")
)

func main() {
	flag.Parse()
	var l int
	for i := 0; i < *count; i++ {
		l = *nameLen
		if *minNameLen != *nameLen {
			l = rand.Intn(*nameLen-*minNameLen) + *minNameLen
		}

		name, err := src.GetRandomName(l)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(name)
	}
}
