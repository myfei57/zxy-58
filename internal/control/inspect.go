package control

type InspectReport struct {
	FlowRate    float64 `json:"flow_rate"`
	ActivePump  string  `json:"active_pump"`
	FilterState string  `json:"filter_state"`
	DrainState  string  `json:"drain_state"`
	PHState     string  `json:"ph_state"`
	Residual    float64 `json:"residual"`
	Speed       float64 `json:"speed"`
}

func (c *Controller) Inspect() InspectReport {
	active := c.Roster.Active()
	activeID := ""
	if active != nil {
		activeID = active.ID
	}
	filterState := "filtering"
	if !c.FilterValve.InFilter {
		filterState = "rinsing"
	}
	drainState := "closed"
	if c.DrainValve.Open {
		drainState = "open"
	}
	phState := "stable"
	if c.PH.Current < c.PH.Target {
		phState = "low"
	}
	residual := c.Doser.Amount() / (c.Circuit.Flow.Rate() + 1)
	return InspectReport{
		FlowRate:    c.Circuit.Flow.Rate(),
		ActivePump:  activeID,
		FilterState: filterState,
		DrainState:  drainState,
		PHState:     phState,
		Residual:    residual,
		Speed:       c.Speed.Current,
	}
}
