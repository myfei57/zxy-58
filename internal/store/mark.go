package store

type MarkRecord struct {
	Key      string `json:"key"`
	Executed bool   `json:"executed"`
}

func (s *Store) MarkExecuted(key string) error {
	return s.write("marks", key, MarkRecord{Key: key, Executed: true})
}

func (s *Store) IsExecuted(key string) (bool, error) {
	var rec MarkRecord
	if err := s.read("marks", key, &rec); err != nil {
		return false, err
	}
	return rec.Executed, nil
}
