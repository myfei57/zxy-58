package pool

type Probe struct {
	ID    string
	Zone  string
	Label string
}

func (p *Pool) AddProbe(probeID, zoneID, label string) {
	p.probes[probeID] = Probe{ID: probeID, Zone: zoneID, Label: label}
}

func (p *Pool) ProbeZone(probeID string) (string, bool) {
	pr, ok := p.probes[probeID]
	if !ok {
		return "", false
	}
	return pr.Zone, true
}

func (p *Pool) ProbeIDs() []string {
	out := make([]string, 0, len(p.probes))
	for id := range p.probes {
		out = append(out, id)
	}
	return out
}

func (p *Pool) Probe(probeID string) (Probe, bool) {
	pr, ok := p.probes[probeID]
	return pr, ok
}
