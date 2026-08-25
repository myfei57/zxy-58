package filter

import "poolops/internal/store"

func Retry(s *store.Store, key string, run func() error) error {
	if err := run(); err != nil {
		return err
	}
	return s.MarkExecuted(key)
}
