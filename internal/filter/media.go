package filter

type Media struct {
	Type      string
	Installed int
	Lifetime  int
}

func NewMedia(kind string, lifetime int) Media {
	return Media{Type: kind, Lifetime: lifetime}
}

func (m Media) RemainingLife() int {
	if m.Installed > m.Lifetime {
		return 0
	}
	return m.Lifetime - m.Installed
}

func (m *Media) Age() int {
	m.Installed++
	return m.Installed
}
