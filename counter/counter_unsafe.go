package counter

type counterUnsafe struct {
	delegate counter
}

func (c *counterUnsafe) Increment() {
	c.delegate.Increment()
}

func (c *counterUnsafe) Get() uint {
	return c.delegate.Get()
}

func NewUnsafeCounter() Counter {
	return &counterUnsafe{delegate: newCounter()}
}
