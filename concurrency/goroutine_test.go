package concurrency

import (
	"errors"
	"testing"
)

func TestGo(t *testing.T) {
	Go(func() {
		panic(errors.New("test"))
	})
}
