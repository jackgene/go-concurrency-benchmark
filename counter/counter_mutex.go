package counter

import "sync"

type counterMutex struct {
	mutex    sync.Mutex
	delegate counter
}

func (c *counterMutex) Increment() {
	c.mutex.Lock()
	c.delegate.Increment()
	c.mutex.Unlock()
}

func (c *counterMutex) Get() uint {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.delegate.Get()
}

func NewMutexCounter() Counter {
	return &counterMutex{
		delegate: newCounter(),
	}
}
