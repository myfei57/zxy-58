package store

import "os"

func (s *Store) Exists(kind, key string) bool {
	_, err := os.Stat(s.path(kind, key))
	return err == nil
}

func (s *Store) CountRecords(kind string) int {
	records, err := s.ListRecords(kind)
	if err != nil {
		return 0
	}
	return len(records)
}

func (s *Store) Keys(kind string) []string {
	path := s.Dir + "/" + kind
	entries, err := os.ReadDir(path)
	if err != nil {
		return []string{}
	}
	out := make([]string, 0, len(entries))
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		if len(name) > 5 && name[len(name)-5:] == ".json" {
			out = append(out, name[:len(name)-5])
		}
	}
	return out
}
