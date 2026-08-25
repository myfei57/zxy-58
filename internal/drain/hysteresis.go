package drain

type Hysteresis struct {
	Target  float64
	Band    float64
	filling bool
}

func NewHysteresis(target, band float64) *Hysteresis {
	return &Hysteresis{Target: target, Band: band}
}

func (h *Hysteresis) ShouldFill(level float64) bool {
	if h.filling {
		if level < h.Target {
			return true
		}
		h.filling = false
		return false
	}
	if level < h.Target-h.Band {
		h.filling = true
		return true
	}
	return false
}

func (h *Hysteresis) Filling() bool {
	return h.filling
}
