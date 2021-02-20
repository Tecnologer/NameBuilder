package test

import (
	"testing"

	randname "github.com/tecnologer/NameBuilder/src"
)

func BenchmarkGetRandom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		randname.GetRandom(10)
	}
}
