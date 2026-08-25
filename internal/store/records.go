package store

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Record struct {
	ID     string  `json:"id"`
	Kind   string  `json:"kind"`
	Amount float64 `json:"amount"`
}

func (s *Store) AppendRecord(kind string, rec Record) error {
	if err := s.ensureDir("records"); err != nil {
		return err
	}
	path := s.path("records", kind)
	var list []Record
	if data, err := os.ReadFile(path); err == nil {
		_ = json.Unmarshal(data, &list)
	}
	list = append(list, rec)
	data, err := json.Marshal(list)
	if err != nil {
		return err
	}
	tmp := path + ".tmp"
	if err := os.WriteFile(tmp, data, 0o644); err != nil {
		return err
	}
	return os.Rename(tmp, path)
}

func (s *Store) ListRecords(kind string) ([]Record, error) {
	data, err := os.ReadFile(filepath.Join(s.Dir, "records", kind+".json"))
	if err != nil {
		return nil, err
	}
	var list []Record
	if err := json.Unmarshal(data, &list); err != nil {
		return nil, err
	}
	return list, nil
}
