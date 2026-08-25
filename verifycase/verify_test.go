package verifycase

import (
	"testing"

	"poolops/internal/circulate"
	"poolops/internal/seq"
)

func TestPoReturnValveOrder(t *testing.T) {
	v := circulate.NewReturnValves()
	rec := seq.NewRecorder()
	branches := []string{"shallow", "deep"}
	if err := circulate.Restore(v, branches, rec); err != nil {
		t.Fatalf("restore failed: %v", err)
	}
	if len(rec.Events) != 2 || rec.Events[0] != "branch-open" || rec.Events[1] != "main-open" {
		t.Fatalf("unexpected return valve order: %v", rec.Events)
	}
}
