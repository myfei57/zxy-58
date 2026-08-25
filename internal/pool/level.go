package pool

func (p *Pool) Level(zoneID string) (float64, bool) {
	z, ok := p.zones[zoneID]
	if !ok {
		return 0, false
	}
	return z.Level, true
}

func (p *Pool) SetLevel(zoneID string, level float64) bool {
	z, ok := p.zones[zoneID]
	if !ok {
		return false
	}
	z.Level = level
	return true
}

func (p *Pool) Snapshot() []Zone {
	out := make([]Zone, 0, len(p.order))
	for _, id := range p.order {
		if z, ok := p.zones[id]; ok {
			out = append(out, z.Summary())
		}
	}
	return out
}
