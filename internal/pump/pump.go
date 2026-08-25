package pump

type Pump struct {
	ID      string
	Name    string
	Running bool
}

func NewPump(id, name string) *Pump {
	return &Pump{ID: id, Name: name}
}
