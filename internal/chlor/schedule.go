package chlor

type Window struct {
	Start int
	End   int
}

func NewWindow(start, end int) Window {
	return Window{Start: start, End: end}
}

func (w Window) Covers(hour int) bool {
	if w.Start <= w.End {
		return hour >= w.Start && hour < w.End
	}
	return hour >= w.Start || hour < w.End
}

type DosingSchedule struct {
	Windows []Window
}

func NewDosingSchedule() *DosingSchedule {
	return &DosingSchedule{}
}

func (d *DosingSchedule) Add(w Window) {
	d.Windows = append(d.Windows, w)
}

func (d *DosingSchedule) Active(hour int) bool {
	for _, w := range d.Windows {
		if w.Covers(hour) {
			return true
		}
	}
	return false
}
