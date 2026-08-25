package control

import "poolops/internal/chlor"

type Analysis struct {
	Residual      float64 `json:"residual"`
	ResidualOK    bool    `json:"residual_ok"`
	PressureHigh  bool    `json:"pressure_high"`
	QuotaRemain   uint64  `json:"quota_remain"`
	SpeedHeadroom float64 `json:"speed_headroom"`
}

func (c *Controller) Analyze() Analysis {
	residual := chlor.ComputeResidual(c.Doser.Amount(), c.Circuit.Flow.Rate())
	quotaRemain := c.Allocator.Remaining("shallow")
	return Analysis{
		Residual:      residual.Free,
		ResidualOK:    residual.InRange(0.3, 1.0),
		PressureHigh:  c.Pressure.NeedsBackwash(0.8),
		QuotaRemain:   quotaRemain,
		SpeedHeadroom: c.Speed.MaxRate - c.Speed.Current,
	}
}
