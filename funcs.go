package bench_channel

import (
	"math/rand"
	"sync"
)

var stubChan chan (bool) = make(chan (bool), 1)
var lock sync.RWMutex
var stubBool bool
var readPercentage int

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func toRead() bool {
	return randInt(1, 100) < readPercentage
}

func WithChannel() {
	if toRead() {
		_ = <-stubChan
		stubChan <- false
	} else {
		foo := <-stubChan
		foo = true
		stubChan <- foo
	}
}

func WithLock() {
	if toRead() {
		lock.RLock()
		_ = stubBool
		lock.RUnlock()
	} else {
		lock.Lock()
		stubBool = true
		lock.Unlock()
		return
	}
}
