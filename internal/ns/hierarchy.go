package ns

type Hierarchy struct {
	parent map[string]string
	order  []string
}

func NewHierarchy() *Hierarchy {
	return &Hierarchy{parent: map[string]string{}}
}

func (h *Hierarchy) Add(zoneID, parentID string) {
	if _, exists := h.parent[zoneID]; !exists {
		h.order = append(h.order, zoneID)
	}
	h.parent[zoneID] = parentID
}

func (h *Hierarchy) Parent(zoneID string) (string, bool) {
	parent, ok := h.parent[zoneID]
	return parent, ok
}
