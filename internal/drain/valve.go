package drain

type Valve struct {
	Open bool
}

func NewValve() *Valve {
	return &Valve{}
}

func (v *Valve) OpenDrain() {
	v.Open = true
}

func (v *Valve) CloseDrain() {
	v.Open = false
}
