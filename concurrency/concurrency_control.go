package concurrency

import (
	"runtime"
	"sync"
)

type ConcurrencyControl struct {
	Ch chan struct{}
	Wg sync.WaitGroup
}

func NewConcurrencyControl(coefficient float64) *ConcurrencyControl {
	cpuNum := runtime.NumCPU()
	if cpuNum <= 0 {
		cpuNum = 1
	}
	tokenNum := int(coefficient * float64(cpuNum))
	if tokenNum <= 0 {
		tokenNum = 1
	}
	ch := make(chan struct{}, tokenNum)
	for i := 0; i < tokenNum; i++ {
		ch <- struct{}{}
	}
	return &ConcurrencyControl{
		Ch: ch,
	}
}

func (c *ConcurrencyControl) Get() {
	<-c.Ch
	c.Wg.Add(1)
}

func (c *ConcurrencyControl) Put() {
	c.Ch <- struct{}{}
	c.Wg.Done()
}

func (c *ConcurrencyControl) Wait() {
	c.Wg.Wait()
}
