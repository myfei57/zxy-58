package pool

import "github.com/google/uuid"

type Pool struct {
	ID     string
	Name   string
	zones  map[string]*Zone
	order  []string
	probes map[string]Probe
}

func NewPool(name string) *Pool {
	return &Pool{
		ID:     uuid.NewString(),
		Name:   name,
		zones:  map[string]*Zone{},
		probes: map[string]Probe{},
	}
}

func (p *Pool) AddZone(id, name string) *Zone {
	if existing, ok := p.zones[id]; ok {
		return existing
	}
	z := NewZone(id, name)
	p.zones[id] = z
	p.order = append(p.order, id)
	return z
}

func (p *Pool) Zone(id string) (*Zone, bool) {
	z, ok := p.zones[id]
	return z, ok
}

func (p *Pool) ZoneIDs() []string {
	out := make([]string, len(p.order))
	copy(out, p.order)
	return out
}
