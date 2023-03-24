package concurrency

import (
	"fmt"
	"runtime"
)

func Go(f func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				buf := make([]byte, 1<<16)
				runtime.Stack(buf, true)
				fmt.Println("err:", err)
				fmt.Printf("stack trace:\n%v\n", string(buf))
			}
		}()
		f()
	}()
}
