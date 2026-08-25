package verifycase

import (
	"testing"

	"poolops/internal/chlor"
	"poolops/internal/circulate"
)

func TestPoChlorDoseFreshFlow(t *testing.T) {
	flow := circulate.NewFlow(100)
	doser := chlor.NewDoser(flow, 0.5)
	flow.SetRate(200)
	got := doser.Amount()
	if got != 100 {
		t.Fatalf("dose should follow current flow, got %v", got)
	}
}
