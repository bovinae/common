package util

import (
	"fmt"
	"testing"
)

func TestCall(t *testing.T) {
	cs := NewConcurrencySsh(&SshConfig{
		Hosts:    []string{"192.168.13.140", "192.168.13.138"},
		User:     "root",
		Password: "PoD2020@sics",
		Port:     22,
		Timeout:  60,
	})
	respMap := cs.Call("ls -al")
	respMap.Range(func(key, value any) bool {
		fmt.Println(key, ":", value)
		return true
	})
}
