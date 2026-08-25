package drain

import (
	"poolops/internal/seq"
	"poolops/internal/store"
)

func drainDoneKey(poolID, zoneID string) string {
	return poolID + "-" + zoneID
}

func Run(s *store.Store, poolID, zoneID string, level float64, rec *seq.Recorder) error {
	if err := s.WriteLevel(poolID, zoneID, level); err != nil {
		return err
	}
	rec.Add("level-persist")
	if err := s.MarkExecuted(drainDoneKey(poolID, zoneID)); err != nil {
		return err
	}
	rec.Add("drain-done")
	return nil
}
