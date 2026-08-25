package control

import "poolops/internal/seq"

type TreatmentResult struct {
	Active  bool    `json:"active"`
	Zones   int     `json:"zones"`
	Dose    float64 `json:"dose"`
	PH      float64 `json:"ph"`
	Records int     `json:"records"`
}

func (c *Controller) FullTreat(hour int) TreatmentResult {
	result := TreatmentResult{}
	if !c.Schedule.Active(hour) {
		return result
	}
	rec := seq.NewRecorder()
	c.runTreatment(rec)
	result.Active = true
	result.Zones = len(c.Pool.ZoneIDs())
	result.Dose = c.Doser.Amount()
	result.PH = c.PH.Current
	for _, zoneID := range c.Pool.ZoneIDs() {
		target, ok := c.Targets.Get(zoneID)
		if !ok {
			continue
		}
		_ = c.Ledger.Record("dose-"+zoneID, target*c.Circuit.Flow.Rate())
		result.Records++
	}
	_ = c.Store.AppendHistory("treatment", rec.Last())
	return result
}

func (c *Controller) PerZoneDose() map[string]float64 {
	out := map[string]float64{}
	for _, zoneID := range c.Pool.ZoneIDs() {
		target, ok := c.Targets.Get(zoneID)
		if ok {
			out[zoneID] = target * c.Circuit.Flow.Rate()
		}
	}
	return out
}

func (c *Controller) TotalDose() float64 {
	sum := 0.0
	for _, dose := range c.PerZoneDose() {
		sum += dose
	}
	return sum
}
