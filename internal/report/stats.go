package report

type PoolStats struct {
	ZoneCount      int     `json:"zone_count"`
	TotalOccupancy int     `json:"total_occupancy"`
	TotalCapacity  int     `json:"total_capacity"`
	AvgTemperature float64 `json:"avg_temperature"`
}

type FlowStats struct {
	Current  float64 `json:"current"`
	Running  bool    `json:"running"`
	Headroom float64 `json:"headroom"`
}

func (c *Collector) PoolStats() PoolStats {
	stats := PoolStats{}
	var tempSum float64
	for _, z := range c.Pool.Snapshot() {
		stats.ZoneCount++
		stats.TotalOccupancy += z.Occupancy
		stats.TotalCapacity += z.Capacity
		tempSum += z.Temperature
	}
	if stats.ZoneCount > 0 {
		stats.AvgTemperature = tempSum / float64(stats.ZoneCount)
	}
	return stats
}

func (c *Collector) FlowStats() FlowStats {
	return FlowStats{
		Current:  c.Circuit.Rate(),
		Running:  c.Circuit.Flow.Running(),
		Headroom: c.Speed.MaxRate - c.Speed.Current,
	}
}
