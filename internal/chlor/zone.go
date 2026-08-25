package chlor

import "poolops/internal/pool"

type ZoneResolver struct {
	cache map[string]string
}

func NewZoneResolver() *ZoneResolver {
	return &ZoneResolver{cache: map[string]string{}}
}

func (r *ZoneResolver) Resolve(p *pool.Pool, probeID string) string {
	if z, ok := r.cache[probeID]; ok {
		return z
	}
	z, ok := p.ProbeZone(probeID)
	if !ok {
		return ""
	}
	r.cache[probeID] = z
	return z
}
