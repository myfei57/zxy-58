package verifycase

import (
	"testing"

	"poolops/internal/drain"
)

func TestPoFillHysteresisBand(t *testing.T) {
	h := drain.NewHysteresis(100, 5)
	_ = h.ShouldFill(90)
	_ = h.ShouldFill(100)
	if h.ShouldFill(99) {
		t.Fatalf("small dip within band should not restart fill")
	}
}
