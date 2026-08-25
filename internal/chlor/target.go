package chlor

type ZoneTarget struct {
	Zone   string
	Target float64
}

type TargetSet struct {
	items map[string]float64
}

func NewTargetSet() *TargetSet {
	return &TargetSet{items: map[string]float64{}}
}

func (t *TargetSet) Set(zone string, target float64) {
	t.items[zone] = target
}

func (t *TargetSet) Get(zone string) (float64, bool) {
	v, ok := t.items[zone]
	return v, ok
}

func (t *TargetSet) List() []ZoneTarget {
	out := make([]ZoneTarget, 0, len(t.items))
	for zone, target := range t.items {
		out = append(out, ZoneTarget{Zone: zone, Target: target})
	}
	return out
}
