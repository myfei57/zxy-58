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
	dv.OpenDrain()
	rec.Add("drain-open")
	v.SwitchToRinse()
	rec.Add("valve-rinse")
	return nil
}
