package counter

type counter struct {
	count uint
}

func (c *counter) Increment() {
	c.count++
}

func (c *counter) Get() uint {
	return c.count
}

func newCounter() counter {
	return counter{0}
}
