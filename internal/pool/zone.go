package pool

type Zone struct {
	ID          string
	Name        string
	Level       float64
	Temperature float64
	Capacity    int
	Occupancy   int
}

func NewZone(id, name string) *Zone {
	return &Zone{ID: id, Name: name}
}

func (z *Zone) Summary() Zone {
	return Zone{
		ID:          z.ID,
		Name:        z.Name,
		Level:       z.Level,
		Temperature: z.Temperature,
		Capacity:    z.Capacity,
		Occupancy:   z.Occupancy,
	}
}
