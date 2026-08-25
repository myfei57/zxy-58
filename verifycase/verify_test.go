package verifycase

import (
	"testing"

	"poolops/internal/chlor"
	"poolops/internal/circulate"
	"poolops/internal/turb"
)

func TestPoTurbidityFilterOrder(t *testing.T) {
	f := turb.NewFilter(0.2)
	f.Apply(10)
	sample := turb.Sample{Raw: 500, Filtered: f.Apply(500)}
	doser := chlor.NewDoser(circulate.NewFlow(100), 1.0)
	dose := doser.TurbidityDose(sample, 200)
	if dose != 0 {
		t.Fatalf("filtered turbidity should not trigger dose, got %v", dose)
	}
}
