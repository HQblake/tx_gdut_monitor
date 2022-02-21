package mysql

import "sync"

type token chan struct{}

type CientLock struct {
	mx    sync.Mutex
	locks map[string]token
}

func (l *CientLock) Lock(id string) {
	l.mx.Lock()
	ch, ok := l.locks[id]
	if !ok {
		ch = make(token, 1)
		l.locks[id] = ch
	}
	l.mx.Unlock()
	ch <- struct{}{}
}

func (l *CientLock) Unlock(id string) {
	l.mx.Lock()
	ch := l.locks[id]
	<-ch
	l.mx.Unlock()
}

func NewLock() *CientLock {
	return &CientLock{
		mx:    sync.Mutex{},
		locks: make(map[string]token),
	}
}
