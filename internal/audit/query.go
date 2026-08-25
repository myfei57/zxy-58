package audit

import "poolops/internal/store"

func (l *Ledger) Query(kind string) []store.Record {
	records, err := l.List()
	if err != nil {
		return []store.Record{}
	}
	out := make([]store.Record, 0)
	for _, rec := range records {
		if kind == "" || rec.Kind == kind {
			out = append(out, rec)
		}
	}
	return out
}

func (l *Ledger) TotalVolume() float64 {
	summary, err := l.Summarize()
	if err != nil {
		return 0
	}
	return summary.Volume
}

func (l *Ledger) ActionCount(kind string) int {
	summary, err := l.Summarize()
	if err != nil {
		return 0
	}
	return summary.CountFor(kind)
}
