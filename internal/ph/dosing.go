package ph

type Dosing struct {
	AcidStep   float64
	AlkaliStep float64
}

func NewDosing(acidStep, alkaliStep float64) Dosing {
	return Dosing{AcidStep: acidStep, AlkaliStep: alkaliStep}
}

func (d Dosing) Correction(target, current float64) float64 {
	delta := target - current
	if delta < 0 {
		return -delta / d.AcidStep
	}
	return delta / d.AlkaliStep
}

func (d Dosing) Direction(target, current float64) string {
	if current > target {
		return "acid"
	}
	if current < target {
		return "alkali"
	}
	return "hold"
}
