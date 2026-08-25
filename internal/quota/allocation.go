package quota

type Allocation struct {
	Zone  string
	Limit uint64
	used  uint64
}

type Allocator struct {
	items map[string]*Allocation
}

func NewAllocator() *Allocator {
	return &Allocator{items: map[string]*Allocation{}}
}

func (a *Allocator) Set(zone string, limit uint64) {
	a.items[zone] = &Allocation{Zone: zone, Limit: limit}
}

func (a *Allocator) Consume(zone string, amount uint64) uint64 {
	item, ok := a.items[zone]
	if !ok {
		return 0
	}
	item.used += amount
	return item.used
}

func (a *Allocator) Remaining(zone string) uint64 {
	item, ok := a.items[zone]
	if !ok {
		return 0
	}
	if item.used >= item.Limit {
		return 0
	}
	return item.Limit - item.used
}

func (a *Allocator) List() []Allocation {
	out := make([]Allocation, 0, len(a.items))
	for _, item := range a.items {
		out = append(out, *item)
	}
	return out
}
