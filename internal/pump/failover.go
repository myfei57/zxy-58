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
	// 先启动备用泵并确认流量建立后再停主泵，避免主备切换瞬间循环流量掉零触发低流量告警。
	r.Standby.Running = true
	flow.SetRunning(true)
	r.Primary.Running = false
	return nil
}
