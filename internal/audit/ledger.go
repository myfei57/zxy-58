package audit

import "poolops/internal/store"

type Ledger struct {
	store *store.Store
}

func NewLedger(s *store.Store) *Ledger {
	return &Ledger{store: s}
}

func (l *Ledger) Record(action string, amount float64) error {
	entry := NewEntry(action, amount)
	return l.store.AppendRecord("audit", store.Record{ID: entry.ID, Kind: action, Amount: amount})
}

func (l *Ledger) List() ([]store.Record, error) {
	return l.store.ListRecords("audit")
}
