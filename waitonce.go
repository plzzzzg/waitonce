package waitonce

import (
	"sync"
	"time"
)

var allWaitOnce = sync.Map{}

type WaitOnce struct {
	Name       string
	waitOnce   sync.Once
	notifyOnce sync.Once
	notifyChan chan struct{}
}

func GetOrCreate(name string) *WaitOnce {
	loaded, ok := allWaitOnce.Load(name)
	if ok {
		return loaded.(*WaitOnce)
	}
	wo := &WaitOnce{
		Name:       name,
		waitOnce:   sync.Once{},
		notifyOnce: sync.Once{},
		notifyChan: make(chan struct{}, 1),
	}
	allWaitOnce.LoadOrStore(name, wo)
	return wo
}

func (wo *WaitOnce) Wait(t time.Duration) (timeout bool) {
	wo.waitOnce.Do(func() {
		select {
		case <-wo.notifyChan:
		case <-time.After(t):
			timeout = true
		}
	})
	return
}

func (wo *WaitOnce) Done() {
	wo.notifyOnce.Do(func() {
		wo.notifyChan <- struct{}{}
	})
}
