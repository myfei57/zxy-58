package control

import "poolops/internal/seq"

type CycleResult struct {
	Hour     int     `json:"hour"`
	Active   bool    `json:"active"`
	Dose     float64 `json:"dose"`
	Speed    float64 `json:"speed"`
	Backwash bool    `json:"backwash"`
}

func (c *Controller) RunCycle(hour int) CycleResult {
	result := CycleResult{Hour: hour}
	if c.Schedule.Active(hour) {
		rec := seq.NewRecorder()
		c.runTreatment(rec)
		result.Active = true
		result.Dose = c.Doser.Amount()
		result.Speed = c.Speed.Current
		_ = c.Store.AppendHistory("treatment", rec.Last())
	}
	if c.Pressure.NeedsBackwash(0.8) {
		_ = c.Backwash()
		result.Backwash = true
	}
	return result
}
