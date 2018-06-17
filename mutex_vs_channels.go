package mutexvschannels

import "sync"

type accWithMutex struct {
	i  int
	mu *sync.Mutex
}

func (v *accWithMutex) lock()   { v.mu.Lock() }
func (v *accWithMutex) unlock() { v.mu.Unlock() }

type accWithChan struct {
	i  int
	ch chan struct{}
}

func (v *accWithChan) lock()   { <-v.ch }
func (v *accWithChan) unlock() { v.ch <- struct{}{} }

// SeriesSumWithMutex function is only for the sync.Mutex benchmarking.
// Do not use the SeriesSumWithMutex function in production environment
// in any circumstances.
func SeriesSumWithMutex(n int) int {
	var (
		v  = &accWithMutex{mu: new(sync.Mutex)}
		wg sync.WaitGroup
	)
	for i := 0; i <= n; i++ {
		wg.Add(1)
		go func(i int) {
			v.lock()
			defer v.unlock()
			defer wg.Done()
			v.i += i
		}(i)
	}
	wg.Wait()
	return v.i
}

// SeriesSumWithChan function is only for the chan benchmarking.
// Do not use the SeriesSumWithChan function in production environment
// in any circumstances.
func SeriesSumWithChan(n int) int {
	v := &accWithChan{ch: make(chan struct{}, 1)}
	v.ch <- struct{}{}
	var wg sync.WaitGroup
	for i := 0; i <= n; i++ {
		wg.Add(1)
		go func(i int) {
			v.lock()
			defer v.unlock()
			defer wg.Done()
			v.i += i
		}(i)
	}
	wg.Wait()
	return v.i
}
