package control

import "poolops/internal/pump"

type PumpStatus struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Running bool   `json:"running"`
	Active  bool   `json:"active"`
}

func (c *Controller) StartPump(id string) bool {
	p := c.pumpByID(id)
	if p == nil {
		return false
	}
	p.Running = true
	return true
}

func (c *Controller) StopPump(id string) bool {
	p := c.pumpByID(id)
	if p == nil {
		return false
	}
	p.Running = false
	return true
}

func (c *Controller) pumpByID(id string) *pump.Pump {
	if c.Roster.Primary != nil && c.Roster.Primary.ID == id {
		return c.Roster.Primary
	}
	if c.Roster.Standby != nil && c.Roster.Standby.ID == id {
		return c.Roster.Standby
	}
	return nil
}

func (c *Controller) PumpStatuses() []PumpStatus {
	active := c.Roster.Active()
	activeID := ""
	if active != nil {
		activeID = active.ID
	}
	out := make([]PumpStatus, 0, 2)
	for _, p := range []*pump.Pump{c.Roster.Primary, c.Roster.Standby} {
		if p == nil {
			continue
		}
		out = append(out, PumpStatus{ID: p.ID, Name: p.Name, Running: p.Running, Active: p.ID == activeID})
	}
	return out
}
