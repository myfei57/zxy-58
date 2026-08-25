package audit

type Summary struct {
	Total   int
	Actions map[string]int
	Volume  float64
}

func (l *Ledger) Summarize() (Summary, error) {
	records, err := l.List()
	if err != nil {
		return Summary{}, err
	}
	summary := Summary{Actions: map[string]int{}}
	for _, rec := range records {
		summary.Total++
		summary.Volume += rec.Amount
		summary.Actions[rec.Kind]++
	}
	return summary, nil
}

func (s Summary) CountFor(action string) int {
	return s.Actions[action]
}
