package control

import (
	"poolops/internal/audit"
	"poolops/internal/circulate"
	"poolops/internal/drain"
	"poolops/internal/filter"
	"poolops/internal/pool"
	"poolops/internal/pump"
	"poolops/internal/quota"
	"poolops/internal/seq"
	"poolops/internal/store"
)

func filterRetry(s *store.Store, key string, run func() error) error {
	return filter.Retry(s, key, run)
}

func filterBackwash(v *filter.Valve, dv *drain.Valve, rec *seq.Recorder) error {
	return filter.Backwash(v, dv, rec)
}

func drainFill(p *pool.Pool, zoneID string, h *drain.Hysteresis, m *quota.Meter, l *audit.Ledger) (drain.FillResult, error) {
	return drain.Fill(p, zoneID, h, m, l)
}

func drainRun(s *store.Store, poolID, zoneID string, level float64, rec *seq.Recorder) error {
	return drain.Run(s, poolID, zoneID, level, rec)
}

func pumpFailover(r *pump.Roster, f *circulate.Flow) error {
	return pump.Failover(r, f)
}

func circulateRestore(v *circulate.ReturnValves, branches []string, rec *seq.Recorder) error {
	return circulate.Restore(v, branches, rec)
}
