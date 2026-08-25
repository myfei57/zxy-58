package circulate

type SpeedController struct {
	MinRate   float64
	MaxRate   float64
	Current   float64
	Increment float64
}

func NewSpeedController(minRate, maxRate, current float64) *SpeedController {
	return &SpeedController{MinRate: minRate, MaxRate: maxRate, Current: current, Increment: 10}
}

func (s *SpeedController) Increase() float64 {
	if s.Current+s.Increment > s.MaxRate {
		s.Current = s.MaxRate
		return s.Current
	}
	s.Current += s.Increment
	return s.Current
}

func (s *SpeedController) Decrease() float64 {
	if s.Current-s.Increment < s.MinRate {
		s.Current = s.MinRate
		return s.Current
	}
	s.Current -= s.Increment
	return s.Current
}
