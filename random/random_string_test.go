package random

import (
	"fmt"
	"testing"
)

func TestGetRandomString(t *testing.T) {
	fmt.Println(GetRandomString(8))
}

func BenchmarkGetRandomString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetRandomString(8)
	}
}
