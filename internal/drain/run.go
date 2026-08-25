package drain

import (
	"poolops/internal/seq"
	"poolops/internal/store"
)

func drainDoneKey(poolID, zoneID string) string {
	return poolID + "-" + zoneID
}

func Run(s *store.Store, poolID, zoneID string, level float64, rec *seq.Recorder) error {
	// Persist the new water level first and only mark the drain complete
	// afterwards. Writing the completion marker before the level leaves a
	// window where, after a crash, the drain looks done while the level
	// still holds its pre-drain value — which fools the fill system into
	// running against a stale level. The marker must be the last write so
	// it acts as a true commit point.
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
