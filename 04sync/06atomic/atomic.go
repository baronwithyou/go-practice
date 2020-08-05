package main

import (
	"sync"
	"sync/atomic"
	"time"
)

type counter interface {
	Inc()
	Load() int64
}

type commonCounter struct {
	counter int64
}

func (c *commonCounter) Inc() {
	c.counter++
}

func (c *commonCounter) Load() int64 {
	return c.counter
}

type mutexCounter struct {
	counter int64
	mu      sync.Mutex
}

func (m *mutexCounter) Inc() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.counter++
}

func (m *mutexCounter) Load() int64 {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.counter
}

type atomicCounter struct {
	counter int64
}

func (a *atomicCounter) Inc() {
	atomic.AddInt64(&a.counter, 1)
}

func (a *atomicCounter) Load() int64 {
	return atomic.LoadInt64(&a.counter)
}

func test(c counter, size int) (int64, time.Duration) {
	var wg sync.WaitGroup
	startTime := time.Now()
	for i := 0; i < size; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			c.Inc()
		}()
	}
	wg.Wait()

	return c.Load(), time.Since(startTime)
}
