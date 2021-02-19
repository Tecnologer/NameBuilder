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

	builder "github.com/tecnologer/NameBuilder/src"
)

func main() {
	nameLen := 6
	name, err := builder.GetRandomName(nameLen)
	if err != nil {
		panic(err)
	}

	fmt.Println(name)
}
```