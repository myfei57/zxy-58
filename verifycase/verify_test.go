package verifycase

import (
	"testing"

	"poolops/internal/drain"
	"poolops/internal/filter"
	"poolops/internal/seq"
)

func TestPoFilterBackwashOrder(t *testing.T) {
	v := filter.NewValve()
	dv := drain.NewValve()
	rec := seq.NewRecorder()
	if err := filter.Backwash(v, dv, rec); err != nil {
		t.Fatalf("backwash failed: %v", err)
	}
	if len(rec.Events) != 2 || rec.Events[0] != "drain-open" || rec.Events[1] != "valve-rinse" {
		t.Fatalf("unexpected backwash order: %v", rec.Events)
	}
}
