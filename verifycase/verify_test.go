package verifycase

import (
	"testing"

	"poolops/internal/chlor"
	"poolops/internal/pool"
)

func TestPoSensorZoneMappingFresh(t *testing.T) {
	resolver := chlor.NewZoneResolver()
	p := pool.NewPool("demo")
	p.AddZone("shallow", "浅水区")
	p.AddZone("deep", "深水区")
	p.AddProbe("probe1", "shallow", "浅水区余氯探头")
	_ = resolver.Resolve(p, "probe1")
	p.AddProbe("probe2", "deep", "深水区余氯探头")
	got := resolver.Resolve(p, "probe2")
	if got != "deep" {
		t.Fatalf("probe2 should map to deep zone, got %q", got)
	}
}
