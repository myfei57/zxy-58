package verifycase

import (
	"testing"

	"poolops/internal/filter"
	"poolops/internal/store"
)

func TestPoBackwashNoDuplicate(t *testing.T) {
	st := store.New(t.TempDir())
	count := 0
	run := func() error {
		count++
		return nil
	}
	_ = filter.Retry(st, "backwash-tank-1", run)
	_ = filter.Retry(st, "backwash-tank-1", run)
	if count != 1 {
		t.Fatalf("backwash should run once, ran %d times", count)
	}
}
