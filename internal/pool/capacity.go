package pool

type Capacity struct {
	ZoneID    string
	Capacity  int
	Occupancy int
}

func (p *Pool) SetCapacity(zoneID string, capacity int) bool {
	z, ok := p.zones[zoneID]
	if !ok {
		return false
	}
	z.Capacity = capacity
	return true
}

func (p *Pool) SetOccupancy(zoneID string, occupancy int) bool {
	z, ok := p.zones[zoneID]
	if !ok {
		return false
	}
	z.Occupancy = occupancy
	return true
}

func (p *Pool) Capacities() []Capacity {
	out := make([]Capacity, 0, len(p.order))
	for _, id := range p.order {
		if z, ok := p.zones[id]; ok {
			out = append(out, Capacity{ZoneID: id, Capacity: z.Capacity, Occupancy: z.Occupancy})
		}
	}
	return out
}
