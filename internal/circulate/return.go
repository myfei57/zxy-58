package circulate

import (
	"errors"

	"poolops/internal/seq"
)

var errNilReturn = errors.New("nil return valves")

type ReturnValves struct {
	MainOpen     bool
	BranchesOpen map[string]bool
}

func NewReturnValves() *ReturnValves {
	return &ReturnValves{BranchesOpen: map[string]bool{}}
}

func (r *ReturnValves) OpenBranch(zone string) {
	r.BranchesOpen[zone] = true
}

func (r *ReturnValves) OpenMain() {
	r.MainOpen = true
}

func (r *ReturnValves) BranchOpen(zone string) bool {
	return r.BranchesOpen[zone]
}

func Restore(v *ReturnValves, branches []string, rec *seq.Recorder) error {
	if v == nil {
		return errNilReturn
	}
	v.OpenMain()
	rec.Add("main-open")
	for _, b := range branches {
		v.OpenBranch(b)
	}
	rec.Add("branch-open")
	return nil
}
