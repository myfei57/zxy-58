package filter

type Valve struct {
	InFilter bool
}

func NewValve() *Valve {
	return &Valve{InFilter: true}
}

func (v *Valve) SwitchToRinse() {
	v.InFilter = false
}

func (v *Valve) SwitchToFilter() {
	v.InFilter = true
}
