# Random Name Builder

Generates names ramdomly.

## Usage

### Install

```bash
go get -u github.com/tecnologer/NameBuilder
```

### Example

```golang
import (
	"fmt"

	randname "github.com/tecnologer/NameBuilder/src"
)

func main() {
	nameLen := 5
	seed := "Da"
	//only lenght
	name, err := randname.GetRandom(nameLen)
	if err != nil {
		panic(err)
	}

	fmt.Println(name)
	// Example output: Ciless

	name, err := randname.GetRandomWithSeed(seed, nameLen)
	if err != nil {
		panic(err)
	}

	fmt.Println(name)
	// Example output: Danice
}
```

### CLI

[Donwload][1] the binary for your system or...

1. `cd $GOPATH/src/github.com/tecnologer/NameBuilder/test/cmd`
2. `make`
3. `./name-random`

   ```bash
   Usage of ./name-random:
   -append
    	flag to indicate if append or replace the current data
   -count int
    	how many names should be generated (default 1)
   -len int
    	length of the name(s) (default 6)
   -len-min int
    	the minimum of lenght of the name(s) (default 4)
   -seed string
    	the minimum of lenght of the name(s)
   -source string
    	the source file to load data
   ```

### Example CLI

```bash
$ ./name-random -len 12 -len-min 5 -seed Tecno
> Tecnoeniche

$ ./name-random -len 8 -len-min 5 -seed Tecno -count 5
> Tecnoeasila
> Tecnoriass
> Tecnores
> Tecnondinan
> Tecnoe
```

[1]: https://github.com/Tecnologer/NameBuilder/releases
