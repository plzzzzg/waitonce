package waitonce

import (
	"sync"
	"testing"
	"time"
)

func TestWaitOnce(t *testing.T) {
	waitOnceID := "x"
	wg := sync.WaitGroup{}
	loopSize := 5
	lock := sync.Mutex{}
	l := make([]int64, loopSize+1)
	for i := 0; i < loopSize; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			timeout := GetOrCreateWaitOnce(waitOnceID).Wait(time.Second)
			if timeout {
				return
			}
			lock.Lock()
			l = append(l, time.Now().Unix())
			lock.Unlock()
		}(i)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(time.Millisecond * 500)
		lock.Lock()
		l = append(l, -1)
		lock.Unlock()
		GetOrCreateWaitOnce(waitOnceID).Done()
	}()

	wg.Wait()
	if l[0] != -1 {
		t.Failed()
	}
}

func TestWaitOnceTimeout(t *testing.T) {
	waitOnceID := "x"
	wg := sync.WaitGroup{}
	loopSize := 10
	lock := sync.Mutex{}
	l := make([]int64, loopSize+1)
	for i := 0; i < loopSize; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			timeout := GetOrCreateWaitOnce(waitOnceID).Wait(time.Millisecond * time.Duration(i*100))
			if timeout {
				return
			}
			lock.Lock()
			l = append(l, time.Now().Unix())
			lock.Unlock()
		}(i)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(time.Millisecond * 550)
		lock.Lock()
		l = append(l, -1)
		lock.Unlock()
		GetOrCreateWaitOnce(waitOnceID).Done()
	}()

	wg.Wait()
	if len(l) != 5 {
		t.Failed()
	}
}

func ExampleGetOrCreateWaitOnce() {
	waitOnceID := "preload"

	go func() {
		if timeout := GetOrCreateWaitOnce(waitOnceID).Wait(time.Second); timeout {
			// fallback when fallback
		} else {
			// do something after preloading done
		}
	}()

	// preload async
	go func() {
		// preloading
		time.Sleep(time.Second)
		GetOrCreateWaitOnce(waitOnceID).Done()
	}()
}
