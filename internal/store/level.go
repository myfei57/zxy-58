package store

type LevelRecord struct {
	PoolID string  `json:"pool_id"`
	ZoneID string  `json:"zone_id"`
	Level  float64 `json:"level"`
}

func (s *Store) WriteLevel(poolID, zoneID string, level float64) error {
	return s.write("levels", poolID+"-"+zoneID, LevelRecord{PoolID: poolID, ZoneID: zoneID, Level: level})
}

func (s *Store) ReadLevel(poolID, zoneID string) (float64, error) {
	var rec LevelRecord
	if err := s.read("levels", poolID+"-"+zoneID, &rec); err != nil {
		return 0, err
	}
	return rec.Level, nil
}
