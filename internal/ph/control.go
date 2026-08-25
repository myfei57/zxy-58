package ph

import (
	"poolops/internal/chlor"
	"poolops/internal/seq"
)

type Control struct {
	Target  float64
	Current float64
}

func NewControl(target, current float64) *Control {
	return &Control{Target: target, Current: current}
}

func (c *Control) Stabilize() float64 {
	if c.Current < c.Target {
		c.Current = c.Target
	}
	return c.Current
}

func Run(c *Control, d *chlor.Doser, rec *seq.Recorder) {
	c.Stabilize()
	rec.Add("ph-stabilize")
	d.Apply()
	rec.Add("chlor-dose")
}
