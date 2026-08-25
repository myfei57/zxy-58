package verifycase

import (
	"testing"

	"poolops/internal/drain"
	"poolops/internal/seq"
	"poolops/internal/store"
)

func TestPoDrainLevelOrder(t *testing.T) {
	st := store.New(t.TempDir())
	rec := seq.NewRecorder()
	if err := drain.Run(st, "pool-1", "zone-1", 1.5, rec); err != nil {
		t.Fatalf("drain run failed: %v", err)
	}
	if len(rec.Events) != 2 || rec.Events[0] != "level-persist" || rec.Events[1] != "drain-done" {
		t.Fatalf("unexpected drain order: %v", rec.Events)
	}
}
