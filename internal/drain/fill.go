package drain

import (
	"errors"

	"poolops/internal/audit"
	"poolops/internal/pool"
	"poolops/internal/quota"
)

var errZoneNotFound = errors.New("zone not found")

type FillResult struct {
	Amount float64
}

func Fill(p *pool.Pool, zoneID string, h *Hysteresis, m *quota.Meter, ledger *audit.Ledger) (FillResult, error) {
	level, ok := p.Level(zoneID)
	if !ok {
		return FillResult{}, errZoneNotFound
	}
	if !h.ShouldFill(level) {
		return FillResult{Amount: 0}, nil
	}
	amount := h.Target - level
	m.Record(uint64(amount * 1000))
	if ledger != nil {
		_ = ledger.Record("fill", amount)
	}
	return FillResult{Amount: amount}, nil
}
