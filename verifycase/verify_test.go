package verifycase

import (
	"testing"

	"poolops/internal/chlor"
	"poolops/internal/circulate"
	"poolops/internal/ph"
	"poolops/internal/seq"
)

func TestPoPhChlorOrder(t *testing.T) {
	flow := circulate.NewFlow(100)
	doser := chlor.NewDoser(flow, 0.5)
	ctrl := ph.NewControl(7.4, 7.1)
	rec := seq.NewRecorder()
	ph.Run(ctrl, doser, rec)
	if len(rec.Events) != 2 || rec.Events[0] != "ph-stabilize" || rec.Events[1] != "chlor-dose" {
		t.Fatalf("unexpected ph-chlor order: %v", rec.Events)
	}
}
