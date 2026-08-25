package report

type ZoneDetailReport struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Level       float64 `json:"level"`
	Temperature float64 `json:"temperature"`
	Capacity    int     `json:"capacity"`
	Occupancy   int     `json:"occupancy"`
	Load        float64 `json:"load"`
}

func (c *Collector) ZoneDetails() []ZoneDetailReport {
	out := make([]ZoneDetailReport, 0, len(c.Pool.ZoneIDs()))
	for _, z := range c.Pool.Snapshot() {
		load := 0.0
		if z.Capacity > 0 {
			load = float64(z.Occupancy) / float64(z.Capacity)
		}
		out = append(out, ZoneDetailReport{
			ID:          z.ID,
			Name:        z.Name,
			Level:       z.Level,
			Temperature: z.Temperature,
			Capacity:    z.Capacity,
			Occupancy:   z.Occupancy,
			Load:        load,
		})
	}
	return out
}
