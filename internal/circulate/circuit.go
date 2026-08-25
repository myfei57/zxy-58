package circulate

type Circuit struct {
	Flow    *Flow
	Running bool
}

func NewCircuit(rate float64) *Circuit {
	return &Circuit{Flow: NewFlow(rate), Running: rate > 0}
}

func (c *Circuit) Rate() float64 {
	return c.Flow.Rate()
}
