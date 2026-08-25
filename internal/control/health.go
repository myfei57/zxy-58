package control

type HealthStatus struct {
	Status     string            `json:"status"`
	Components map[string]string `json:"components"`
	IssueCount int               `json:"issue_count"`
}

func (c *Controller) Health() HealthStatus {
	components := map[string]string{
		"pool":        "ok",
		"circulation": "ok",
		"pump":        "ok",
		"filter":      "ok",
		"chlor":       "ok",
		"drain":       "ok",
	}
	issues := 0
	if c.Circuit.Flow.Rate() <= 0 {
		components["circulation"] = "no-flow"
		issues++
	}
	if c.Roster.Active() == nil {
		components["pump"] = "no-active-pump"
		issues++
	}
	if c.Media.RemainingLife() <= 0 {
		components["filter"] = "media-expired"
		issues++
	}
	if c.Doser.Amount() < 0.3 {
		components["chlor"] = "residual-low"
		issues++
	}
	if c.Hysteresis.Filling() {
		components["drain"] = "filling"
	}
	status := "ok"
	if issues > 0 {
		status = "degraded"
	}
	return HealthStatus{Status: status, Components: components, IssueCount: issues}
}

func (c *Controller) FlowHealthy() bool {
	return c.Circuit.Flow.Running() && c.Circuit.Flow.Rate() > 0
}
