package control

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

type Controller struct {
	Pool         *pool.Pool
	Circuit      *circulate.Circuit
	Speed        *circulate.SpeedController
	ReturnValves *circulate.ReturnValves
	Roster       *pump.Roster
	PrimaryTel   *pump.Telemetry
	StandbyTel   *pump.Telemetry
	FilterValve  *filter.Valve
	FilterTank   *filter.Tank
	Media        *filter.Media
	Pressure     *filter.Pressure
	DrainValve   *drain.Valve
	Doser        *chlor.Doser
	Targets      *chlor.TargetSet
	Schedule     *chlor.DosingSchedule
	PH           *ph.Control
	PHDosing     *ph.Dosing
	TurbFilter   *turb.Filter
	Sensor       *turb.Sensor
	Hysteresis   *drain.Hysteresis
	Quota        *quota.Meter
	Allocator    *quota.Allocator
	Ledger       *audit.Ledger
	Store        *store.Store
	Resolver     *chlor.ZoneResolver
	Hierarchy    *ns.Hierarchy
}
