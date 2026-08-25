package report

import (
	"poolops/internal/audit"
	"poolops/internal/chlor"
	"poolops/internal/circulate"
	"poolops/internal/drain"
	"poolops/internal/filter"
	"poolops/internal/ns"
	"poolops/internal/ph"
	"poolops/internal/pool"
	"poolops/internal/pump"
	"poolops/internal/quota"
	"poolops/internal/store"
	"poolops/internal/turb"
)

type Collector struct {
	Pool        *pool.Pool
	Circuit     *circulate.Circuit
	Speed       *circulate.SpeedController
	Return      *circulate.ReturnValves
	Roster      *pump.Roster
	PrimaryTel  *pump.Telemetry
	StandbyTel  *pump.Telemetry
	FilterValve *filter.Valve
	FilterTank  *filter.Tank
	Media       *filter.Media
	Pressure    *filter.Pressure
	DrainValve  *drain.Valve
	Doser       *chlor.Doser
	Targets     *chlor.TargetSet
	PH          *ph.Control
	PHDosing    *ph.Dosing
	TurbFilter  *turb.Filter
	Sensor      *turb.Sensor
	Hysteresis  *drain.Hysteresis
	Quota       *quota.Meter
	Allocator   *quota.Allocator
	Ledger      *audit.Ledger
	Hierarchy   *ns.Hierarchy
	Store       *store.Store
}

func (c *Collector) Collect() Snapshot {
	return Snapshot{
		Pool:        c.collectPool(),
		Circulation: c.collectCirculation(),
		Filtration:  c.collectFiltration(),
		Pumps:       c.collectPumps(),
		Chemicals:   c.collectChemicals(),
		Turbidity:   c.collectTurbidity(),
		Drain:       c.collectDrain(),
		Quota:       c.collectQuota(),
		Audit:       c.collectAudit(),
		Alarms:      c.collectAlarms(),
	}
}

func (c *Collector) collectPool() PoolReport {
	report := PoolReport{ID: c.Pool.ID, Name: c.Pool.Name}
	for _, z := range c.Pool.Snapshot() {
		parent, _ := c.Hierarchy.Parent(z.ID)
		report.Zones = append(report.Zones, ZoneDetail{
			ID:          z.ID,
			Name:        z.Name,
			Level:       z.Level,
			Temperature: z.Temperature,
			Capacity:    z.Capacity,
			Occupancy:   z.Occupancy,
			Parent:      parent,
		})
	}
	for _, probeID := range c.Pool.ProbeIDs() {
		if pr, ok := c.Pool.Probe(probeID); ok {
			report.Probes = append(report.Probes, ProbeDetail{ID: pr.ID, Zone: pr.Zone})
		}
	}
	return report
}

func (c *Collector) collectCirculation() CirculationReport {
	branches := make([]string, 0)
	for _, id := range c.Pool.ZoneIDs() {
		if c.Return.BranchOpen(id) {
			branches = append(branches, id)
		}
	}
	return CirculationReport{
		FlowRate:     c.Circuit.Rate(),
		Running:      c.Circuit.Flow.Running(),
		Speed:        c.Speed.Current,
		MainOpen:     c.Return.MainOpen,
		BranchesOpen: branches,
	}
}

func (c *Collector) collectFiltration() FiltrationReport {
	return FiltrationReport{
		TankRuntime: c.FilterTank.Runtime,
		InFilter:    c.FilterValve.InFilter,
		Pressure:    c.Pressure.Differential(),
		MediaLife:   c.Media.RemainingLife(),
	}
}

func (c *Collector) collectPumps() PumpsReport {
	active := c.Roster.Active()
	activeID := ""
	if active != nil {
		activeID = active.ID
	}
	return PumpsReport{
		Primary: PumpDetail{ID: c.Roster.Primary.ID, Name: c.Roster.Primary.Name, Running: c.Roster.Primary.Running, Cycles: c.PrimaryTel.Cycles()},
		Standby: PumpDetail{ID: c.Roster.Standby.ID, Name: c.Roster.Standby.Name, Running: c.Roster.Standby.Running, Cycles: c.StandbyTel.Cycles()},
		Active:  activeID,
	}
}

func (c *Collector) collectChemicals() ChemicalsReport {
	return ChemicalsReport{
		ChlorineDose: c.Doser.Amount(),
		Targets:      c.Targets.List(),
		PHTarget:     c.PH.Target,
		PHCurrent:    c.PH.Current,
		PHDirection:  c.PHDosing.Direction(c.PH.Target, c.PH.Current),
	}
}

func (c *Collector) collectTurbidity() TurbidityReport {
	return TurbidityReport{
		Filtered:  c.TurbFilter.Last(),
		Raw:       c.TurbFilter.Raw(),
		SensorMax: c.Sensor.RangeMax,
	}
}

func (c *Collector) collectDrain() DrainReport {
	return DrainReport{
		DrainOpen: c.DrainValve.Open,
		Filling:   c.Hysteresis.Filling(),
		Target:    c.Hysteresis.Target,
		Band:      c.Hysteresis.Band,
	}
}

func (c *Collector) collectQuota() QuotaReport {
	return QuotaReport{
		Total:       c.Quota.Total(),
		Allocations: c.Allocator.List(),
	}
}

func (c *Collector) collectAudit() AuditReport {
	summary, err := c.Ledger.Summarize()
	if err != nil {
		return AuditReport{}
	}
	return AuditReport{Total: summary.Total, Volume: summary.Volume, Actions: summary.Actions}
}

func (c *Collector) collectAlarms() []AlarmItem {
	alarms := make([]AlarmItem, 0)
	for _, z := range c.Pool.Snapshot() {
		if z.Level < 0.2 {
			alarms = append(alarms, AlarmItem{ID: "low-" + z.ID, Message: z.Name + " 水位过低", Level: "critical"})
		}
		if z.Temperature > 30 {
			alarms = append(alarms, AlarmItem{ID: "hot-" + z.ID, Message: z.Name + " 水温过高", Level: "warning"})
		}
	}
	if c.Doser.Amount() < 0.3 {
		alarms = append(alarms, AlarmItem{ID: "residual-low", Message: "余氯偏低", Level: "warning"})
	}
	if c.Pressure.Differential() >= 0.8 {
		alarms = append(alarms, AlarmItem{ID: "pressure-high", Message: "滤罐压差过高", Level: "warning"})
	}
	return alarms
}
