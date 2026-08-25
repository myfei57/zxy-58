package quota

type Meter struct {
	counter *Counter
}

func NewMeter(limit uint64) *Meter {
	return &Meter{counter: NewCounter(limit)}
}

func (m *Meter) Record(amount uint64) uint64 {
	return m.counter.Add(amount)
}

func (m *Meter) Total() uint64 {
	return m.counter.Value()
}
