package concurrency

import (
	"runtime"
	"sync"
)

type ConcurrencyControl struct {
	ch chan struct{}
	wg sync.WaitGroup
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
	return NewControllerWithTokenNum(tokenNum)
}

func NewControllerWithTokenNum(tokenNum int) *ConcurrencyControl {
	ch := make(chan struct{}, tokenNum)
	return &ConcurrencyControl{
		ch: ch,
	}
}

func (c *ConcurrencyControl) Get() {
	c.ch <- struct{}{}
	c.wg.Add(1)
}

func (c *ConcurrencyControl) Put() {
	<-c.ch
	c.wg.Done()
}

func (c *ConcurrencyControl) Wait() {
	c.wg.Wait()
}
