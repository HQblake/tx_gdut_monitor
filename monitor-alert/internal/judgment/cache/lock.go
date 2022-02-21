package cache

import "sync"

type token chan struct{}

type CacheLock struct {
	mx    sync.Mutex
	locks map[string]token
}

func (l *CacheLock) Lock(id string) {
	l.mx.Lock()
	ch, ok := l.locks[id]
	if !ok {
		ch = make(token, 1)
		l.locks[id] = ch
	}
	l.mx.Unlock()
	ch <- struct{}{}
}

func (l *CacheLock) Unlock(id string) {
	l.mx.Lock()
	ch := l.locks[id]
	<-ch
	l.mx.Unlock()
}

func NewLock() *CacheLock {
	return &CacheLock{
		mx:    sync.Mutex{},
		locks: make(map[string]token),
	}
}
