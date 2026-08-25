package turb

type Calibration struct {
	Offset float64
	Slope  float64
}

func NewCalibration(offset, slope float64) Calibration {
	return Calibration{Offset: offset, Slope: slope}
}

func (c Calibration) Calibrate(raw float64) float64 {
	return raw*c.Slope + c.Offset
}
