package report

type TimelineEntry struct {
	At      string `json:"at"`
	Key     string `json:"key"`
	Message string `json:"message"`
}

type Trend struct {
	Key   string `json:"key"`
	Count int    `json:"count"`
	First string `json:"first"`
	Last  string `json:"last"`
}

func (c *Collector) History(key string) []TimelineEntry {
	entries, err := c.Store.History(key)
	if err != nil {
		return []TimelineEntry{}
	}
	out := make([]TimelineEntry, 0, len(entries))
	for _, entry := range entries {
		out = append(out, TimelineEntry{
			At:      entry.At.Format("2006-01-02T15:04:05Z07:00"),
			Key:     entry.Key,
			Message: entry.Message,
		})
	}
	return out
}

func (c *Collector) Trends() []Trend {
	keys := []string{"treatment", "backwash", "drain", "restore"}
	out := make([]Trend, 0, len(keys))
	for _, key := range keys {
		entries, err := c.Store.History(key)
		if err != nil || len(entries) == 0 {
			out = append(out, Trend{Key: key})
			continue
		}
		out = append(out, Trend{
			Key:   key,
			Count: len(entries),
			First: entries[0].At.Format("2006-01-02T15:04:05Z07:00"),
			Last:  entries[len(entries)-1].At.Format("2006-01-02T15:04:05Z07:00"),
		})
	}
	return out
}
