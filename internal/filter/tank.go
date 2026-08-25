package filter

type Tank struct {
	ID      string
	Runtime int
}

func NewTank(id string) *Tank {
	return &Tank{ID: id}
}

func (t *Tank) Rinse() {
	t.Runtime++
}
