package chlor

import (
	"poolops/internal/circulate"
	"poolops/internal/turb"
)

type Doser struct {
	Flow   *circulate.Flow
	Target float64
}

func NewDoser(flow *circulate.Flow, target float64) *Doser {
	return &Doser{Flow: flow, Target: target}
}

func (d *Doser) Amount() float64 {
	return d.Target * d.Flow.Rate()
}

func (d *Doser) Apply() float64 {
	return d.Amount()
}

func (d *Doser) TurbidityDose(s turb.Sample, limit float64) float64 {
	value := s.Filtered
	if value > limit {
		return d.Target * (value - limit)
	}
	return 0
}
