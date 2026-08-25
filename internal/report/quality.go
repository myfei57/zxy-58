package report

import "poolops/internal/chlor"

type QualityPoint struct {
	Zone      string  `json:"zone"`
	Turbidity float64 `json:"turbidity"`
	Residual  float64 `json:"residual"`
	PH        float64 `json:"ph"`
}

func (c *Collector) Quality() []QualityPoint {
	residual := chlor.ComputeResidual(c.Doser.Amount(), c.Circuit.Rate())
	out := make([]QualityPoint, 0, len(c.Pool.ZoneIDs()))
	for _, z := range c.Pool.Snapshot() {
		out = append(out, QualityPoint{
			Zone:      z.Name,
			Turbidity: c.TurbFilter.Last(),
			Residual:  residual.Free,
			PH:        c.PH.Current,
		})
	}
	return out
}

func (c *Collector) QualitySummary() map[string]float64 {
	residual := chlor.ComputeResidual(c.Doser.Amount(), c.Circuit.Rate())
	return map[string]float64{
		"turbidity": c.TurbFilter.Last(),
		"residual":  residual.Free,
		"ph":        c.PH.Current,
		"flow":      c.Circuit.Rate(),
	}
}
