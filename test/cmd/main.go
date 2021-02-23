package main

import (
	"flag"
	"fmt"
	"math/rand"

	randName "github.com/tecnologer/NameBuilder/src"
)

var (
	count      = flag.Int("count", 1, "how many names should be generated")
	nameLen    = flag.Int("len", 4, "length of the name(s)")
	minNameLen = flag.Int("len-min", 4, "the minimum of lenght of the name(s)")
	seed       = flag.String("seed", "", "the minimum of lenght of the name(s)")
	dataSource = flag.String("source", "", "the source file to load data")
	append     = flag.Bool("append", false, "flag to indicate if append or replace the current data")
)

func main() {
	// defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	//defer profile.Start(profile.MemProfile, profile.MemProfileRate(1), profile.ProfilePath(".")).Stop()
	// defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()

	flag.Parse()

	if *dataSource != "" {
		if *append {
			randName.AppendData(*dataSource)
		} else {
			randName.LoadData(*dataSource)
		}
	}

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
