package turb

type Filter struct {
	Alpha float64
	last  float64
	raw   float64
}

func NewFilter(alpha float64) *Filter {
	return &Filter{Alpha: alpha}
}

func (f *Filter) Apply(raw float64) float64 {
	f.raw = raw
	f.last = f.Alpha*raw + (1-f.Alpha)*f.last
	return f.last
}

func (f *Filter) Last() float64 {
	return f.last
}

func (f *Filter) Raw() float64 {
	return f.raw
}
