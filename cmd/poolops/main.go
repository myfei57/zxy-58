package main

import (
	"net/http"
	"os"

	"poolops/internal/audit"
	"poolops/internal/chlor"
	"poolops/internal/circulate"
	"poolops/internal/console"
	"poolops/internal/control"
	"poolops/internal/drain"
	"poolops/internal/filter"
	"poolops/internal/ns"
	"poolops/internal/ph"
	"poolops/internal/pool"
	"poolops/internal/pump"
	"poolops/internal/quota"
	"poolops/internal/report"
	"poolops/internal/store"
	"poolops/internal/turb"
)

func main() {
	dataDir := os.Getenv("POOLOPS_DATA_DIR")
	if dataDir == "" {
		dataDir = "data"
	}
	templateDir := os.Getenv("POOLOPS_TEMPLATE_DIR")
	if templateDir == "" {
		templateDir = "internal/console/templates"
	}
	addr := os.Getenv("POOLOPS_ADDR")
	if addr == "" {
		addr = "127.0.0.1:8080"
	}

	st := store.New(dataDir)
	p := pool.NewPool("训练池")
	p.AddZone("shallow", "浅水区")
	p.AddZone("deep", "深水区")
	p.SetLevel("shallow", 1.2)
	p.SetLevel("deep", 1.8)
	p.SetTemperature("shallow", 26.5)
	p.SetTemperature("deep", 27.0)
	p.SetCapacity("shallow", 120)
	p.SetCapacity("deep", 80)
	p.SetOccupancy("shallow", 40)
	p.SetOccupancy("deep", 25)
	p.AddProbe("probe-shallow", "shallow", "浅水区余氯探头")
	p.AddProbe("probe-deep", "deep", "深水区余氯探头")

	hierarchy := ns.NewHierarchy()
	hierarchy.Add("shallow", "")
	hierarchy.Add("deep", "")

	flow := circulate.NewFlow(500)
	circuit := circulate.NewCircuit(500)
	speed := circulate.NewSpeedController(300, 800, 500)
	returns := circulate.NewReturnValves()
	returns.OpenBranch("shallow")
	returns.OpenBranch("deep")
	returns.OpenMain()

	primary := pump.NewPump("pump-1", "主泵")
	primary.Running = true
	standby := pump.NewPump("pump-2", "备泵")
	roster := pump.NewRoster(primary, standby)
	primaryTel := &pump.Telemetry{RuntimeHours: 120, StartCount: 30, StopCount: 29}
	standbyTel := &pump.Telemetry{RuntimeHours: 8, StartCount: 3, StopCount: 2}

	filterValve := filter.NewValve()
	filterTank := filter.NewTank("tank-1")
	filterTank.Runtime = 42
	media := filter.NewMedia("石英砂", 365)
	media.Installed = 100
	pressure := filter.NewPressure(0.45, 0.30)
	drainValve := drain.NewValve()

	doser := chlor.NewDoser(flow, 0.5)
	targets := chlor.NewTargetSet()
	targets.Set("shallow", 0.5)
	targets.Set("deep", 0.5)
	schedule := chlor.NewDosingSchedule()
	schedule.Add(chlor.NewWindow(6, 22))

	phControl := ph.NewControl(7.4, 7.1)
	phDosing := ph.NewDosing(0.2, 0.2)
	turbFilter := turb.NewFilter(0.2)
	sensor := turb.NewSensor("sensor-1", "shallow", 100, turb.NewCalibration(0, 1))
	hysteresis := drain.NewHysteresis(1.8, 0.1)
	quotaMeter := quota.NewMeter(1000000)
	allocator := quota.NewAllocator()
	allocator.Set("shallow", 100000)
	allocator.Set("deep", 80000)
	ledger := audit.NewLedger(st)
	resolver := chlor.NewZoneResolver()

	collector := &report.Collector{
		Pool:        p,
		Circuit:     circuit,
		Speed:       speed,
		Return:      returns,
		Roster:      roster,
		PrimaryTel:  primaryTel,
		StandbyTel:  standbyTel,
		FilterValve: filterValve,
		FilterTank:  filterTank,
		Media:       &media,
		Pressure:    &pressure,
		DrainValve:  drainValve,
		Doser:       doser,
		Targets:     targets,
		PH:          phControl,
		PHDosing:    &phDosing,
		TurbFilter:  turbFilter,
		Sensor:      &sensor,
		Hysteresis:  hysteresis,
		Quota:       quotaMeter,
		Allocator:   allocator,
		Ledger:      ledger,
		Hierarchy:   hierarchy,
		Store:       st,
	}

	controller := &control.Controller{
		Pool:         p,
		Circuit:      circuit,
		Speed:        speed,
		ReturnValves: returns,
		Roster:       roster,
		PrimaryTel:   primaryTel,
		StandbyTel:   standbyTel,
		FilterValve:  filterValve,
		FilterTank:   filterTank,
		Media:        &media,
		Pressure:     &pressure,
		DrainValve:   drainValve,
		Doser:        doser,
		Targets:      targets,
		Schedule:     schedule,
		PH:           phControl,
		PHDosing:     &phDosing,
		TurbFilter:   turbFilter,
		Sensor:       &sensor,
		Hysteresis:   hysteresis,
		Quota:        quotaMeter,
		Allocator:    allocator,
		Ledger:       ledger,
		Store:        st,
		Resolver:     resolver,
		Hierarchy:    hierarchy,
	}

	pages, err := console.NewPages(templateDir)
	if err != nil {
		panic(err)
	}

	server := &console.Server{Report: collector, Control: controller, Pages: pages}
	if err := http.ListenAndServe(addr, server.Router()); err != nil {
		panic(err)
	}
}
