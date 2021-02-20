package main

import (
	"flag"
	"fmt"
	"math/rand"

	"github.com/pkg/profile"
	randName "github.com/tecnologer/NameBuilder/src"
)

var (
	count      = flag.Int("count", 1, "how many names should be generated")
	nameLen    = flag.Int("len", 4, "length of the name(s)")
	minNameLen = flag.Int("len-min", 4, "the minimum of lenght of the name(s)")
	seed       = flag.String("seed", "", "the minimum of lenght of the name(s)")
)

func main() {
	defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()

	flag.Parse()
	var l int
	for i := 0; i < *count; i++ {
		l = *nameLen
		if *minNameLen != *nameLen {
			l = rand.Intn(*nameLen-*minNameLen) + *minNameLen
		}

		name, err := randName.GetRandomWithSeed(*seed, l)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(name)
	}
}
