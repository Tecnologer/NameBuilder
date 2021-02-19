package main

import (
	"fmt"
	"math/rand"

	"github.com/tecnologer/NameBuilder/src"
)

func main() {

	for i := 0; i < 10; i++ {
		name, err := src.GetRandomName(rand.Intn(10-4) + 4)
		if err != nil {
			continue
		}

		fmt.Println(name)
	}
}
