package pump

import (
	"errors"

	"poolops/internal/circulate"
)

var errInvalidRoster = errors.New("invalid roster")

func Failover(r *Roster, flow *circulate.Flow) error {
	if r == nil || r.Primary == nil || r.Standby == nil {
		return errInvalidRoster
	}
	if r.Standby.Running {
		return nil
	}
	r.Standby.Running = true
	flow.SetRunning(true)
	r.Primary.Running = false
	return nil
}
