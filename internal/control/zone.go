package control

func (c *Controller) ZoneDose(zoneID string) float64 {
	target, ok := c.Targets.Get(zoneID)
	if !ok {
		return c.Doser.Amount()
	}
	return target * c.Circuit.Flow.Rate()
}

func (c *Controller) ZoneFillDue(zoneID string) bool {
	level, ok := c.Pool.Level(zoneID)
	if !ok {
		return false
	}
	return level < c.Hysteresis.Target-c.Hysteresis.Band
}
