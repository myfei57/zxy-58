package quota

type Counter struct {
	Limit uint64
	value uint64
}

func NewCounter(limit uint64) *Counter {
	return &Counter{Limit: limit}
}

func (c *Counter) Add(amount uint64) uint64 {
	if c.Limit == 0 {
		c.value += amount
		return c.value
	}
	c.value = (c.value + amount) % c.Limit
	return c.value
}

func (c *Counter) Value() uint64 {
	return c.value
}
