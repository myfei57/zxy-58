package store

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

type HistoryEntry struct {
	At      time.Time `json:"at"`
	Key     string    `json:"key"`
	Message string    `json:"message"`
}

func (s *Store) AppendHistory(key, message string) error {
	if err := s.ensureDir("history"); err != nil {
		return err
	}
	path := filepath.Join(s.Dir, "history", key+".json")
	var list []HistoryEntry
	if data, err := os.ReadFile(path); err == nil {
		_ = json.Unmarshal(data, &list)
	}
	list = append(list, HistoryEntry{At: time.Now().UTC(), Key: key, Message: message})
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

func (s *Store) History(key string) ([]HistoryEntry, error) {
	data, err := os.ReadFile(filepath.Join(s.Dir, "history", key+".json"))
	if err != nil {
		return nil, err
	}
	var list []HistoryEntry
	if err := json.Unmarshal(data, &list); err != nil {
		return nil, err
	}
	return list, nil
}
