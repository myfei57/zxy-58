package verifycase

import (
	"testing"

	"poolops/internal/circulate"
	"poolops/internal/pump"
)

func TestPoCirculationFailoverOrder(t *testing.T) {
	flow := circulate.NewFlow(500)
	primary := pump.NewPump("pump-1", "主泵")
	primary.Running = true
	standby := pump.NewPump("pump-2", "备泵")
	roster := pump.NewRoster(primary, standby)
	if err := pump.Failover(roster, flow); err != nil {
		t.Fatalf("failover failed: %v", err)
	}
	for _, event := range flow.Events() {
		if event == "flow-stop" {
			t.Fatalf("flow should not stop during failover: %v", flow.Events())
		}
	}
}
