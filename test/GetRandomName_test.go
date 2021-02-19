package test

import (
	"testing"

	"github.com/tecnologer/NameBuilder/src"
)

func BenchmarkGetRandomName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src.GetRandomName(10)
	}
}
