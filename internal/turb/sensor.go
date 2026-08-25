package turb

type Sensor struct {
	ID          string
	Zone        string
	RangeMax    float64
	Calibration Calibration
}

func NewSensor(id, zone string, rangeMax float64, calibration Calibration) Sensor {
	return Sensor{ID: id, Zone: zone, RangeMax: rangeMax, Calibration: calibration}
}

func (s Sensor) Read(raw float64) float64 {
	value := s.Calibration.Calibrate(raw)
	if value > s.RangeMax {
		return s.RangeMax
	}
	return value
}
