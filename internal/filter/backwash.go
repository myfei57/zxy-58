package filter

import (
	"errors"

	"poolops/internal/drain"
	"poolops/internal/seq"
)

var errNilValve = errors.New("nil valve")

func Backwash(v *Valve, dv *drain.Valve, rec *seq.Recorder) error {
	if v == nil || dv == nil {
		return errNilValve
	}
	// 排污阀须先于过滤阀切换，避免阀位切换瞬间杂质回流到池体。
	dv.OpenDrain()
	rec.Add("drain-open")
	v.SwitchToRinse()
	rec.Add("valve-rinse")
	return nil
}
