package filter

type Pressure struct {
	Inlet  float64
	Outlet float64
}

func NewPressure(inlet, outlet float64) Pressure {
	return Pressure{Inlet: inlet, Outlet: outlet}
}

func (p Pressure) Differential() float64 {
	return p.Inlet - p.Outlet
}

func (p Pressure) NeedsBackwash(threshold float64) bool {
	return p.Differential() >= threshold
}
