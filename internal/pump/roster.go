package pump

type Roster struct {
	Primary *Pump
	Standby *Pump
}

func NewRoster(primary, standby *Pump) *Roster {
	return &Roster{Primary: primary, Standby: standby}
}

func (r *Roster) Active() *Pump {
	if r.Primary != nil && r.Primary.Running {
		return r.Primary
	}
	if r.Standby != nil && r.Standby.Running {
		return r.Standby
	}
	return nil
}
