package cache

import "sync"

type token chan struct{}

type CacheLock struct {
	mx    sync.Mutex
	locks map[string]token
}

func (l *CacheLock) Lock(id string) {
	l.mx.Lock()
	_, ok := l.locks[id]
	if !ok {
		l.locks[id] = make(token, 1)
	}
	l.mx.Unlock()
	l.locks[id] <- struct{}{}
}

func (l *CacheLock) Unlock(id string) {
	l.mx.Lock()
	<-l.locks[id]
	l.mx.Unlock()
}

func NewLock() *CacheLock {
	return &CacheLock{
		mx:    sync.Mutex{},
		locks: make(map[string]token),
	}
}
