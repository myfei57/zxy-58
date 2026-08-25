package pool

type Temperature struct {
	ZoneID string
	Value  float64
}

func (p *Pool) SetTemperature(zoneID string, value float64) bool {
	z, ok := p.zones[zoneID]
	if !ok {
		return false
	}
	z.Temperature = value
	return true
}

func (p *Pool) Temperature(zoneID string) (float64, bool) {
	z, ok := p.zones[zoneID]
	if !ok {
		return 0, false
	}
	return z.Temperature, true
}

func (p *Pool) Temperatures() []Temperature {
	out := make([]Temperature, 0, len(p.order))
	for _, id := range p.order {
		if z, ok := p.zones[id]; ok {
			out = append(out, Temperature{ZoneID: id, Value: z.Temperature})
		}
	}
	return out
}
