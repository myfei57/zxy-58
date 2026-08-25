package store

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Store struct {
	Dir string
}

func New(dir string) *Store {
	return &Store{Dir: dir}
}

func (s *Store) path(kind, key string) string {
	return filepath.Join(s.Dir, kind, key+".json")
}

func (s *Store) ensureDir(kind string) error {
	return os.MkdirAll(filepath.Join(s.Dir, kind), 0o755)
}

func (s *Store) write(kind, key string, value any) error {
	if err := s.ensureDir(kind); err != nil {
		return err
	}
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	tmp := s.path(kind, key) + ".tmp"
	if err := os.WriteFile(tmp, data, 0o644); err != nil {
		return err
	}
	return os.Rename(tmp, s.path(kind, key))
}

func (s *Store) read(kind, key string, value any) error {
	data, err := os.ReadFile(s.path(kind, key))
	if err != nil {
		return err
	}
	return json.Unmarshal(data, value)
}
