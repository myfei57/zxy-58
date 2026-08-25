package store

type Snapshot struct {
	Levels  int `json:"levels"`
	Marks   int `json:"marks"`
	Records int `json:"records"`
	History int `json:"history"`
}

func (s *Store) Snapshot() Snapshot {
	return Snapshot{
		Levels:  len(s.Keys("levels")),
		Marks:   len(s.Keys("marks")),
		Records: len(s.Keys("records")),
		History: len(s.Keys("history")),
	}
}
