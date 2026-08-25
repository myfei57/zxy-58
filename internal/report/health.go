package report

type HealthSummary struct {
	Status       string    `json:"status"`
	Pool         PoolStats `json:"pool"`
	Flow         FlowStats `json:"flow"`
	ResidualOK   bool      `json:"residual_ok"`
	PressureHigh bool      `json:"pressure_high"`
}

func (c *Collector) HealthSummary() HealthSummary {
	residual := c.Doser.Amount() / (c.Circuit.Rate() + 1)
	status := "ok"
	if !c.Circuit.Flow.Running() || residual < 0.3 {
		status = "degraded"
	}
	return HealthSummary{
		Status:       status,
		Pool:         c.PoolStats(),
		Flow:         c.FlowStats(),
		ResidualOK:   residual >= 0.3,
		PressureHigh: c.Pressure.NeedsBackwash(0.8),
	}
}
