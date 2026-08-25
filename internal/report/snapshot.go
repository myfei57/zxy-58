package report

import (
	"poolops/internal/chlor"
	"poolops/internal/quota"
)

type ZoneDetail struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Level       float64 `json:"level"`
	Temperature float64 `json:"temperature"`
	Capacity    int     `json:"capacity"`
	Occupancy   int     `json:"occupancy"`
	Parent      string  `json:"parent"`
}

type ProbeDetail struct {
	ID   string `json:"id"`
	Zone string `json:"zone"`
}

type PoolReport struct {
	ID     string        `json:"id"`
	Name   string        `json:"name"`
	Zones  []ZoneDetail  `json:"zones"`
	Probes []ProbeDetail `json:"probes"`
}

type CirculationReport struct {
	FlowRate     float64  `json:"flow_rate"`
	Running      bool     `json:"running"`
	Speed        float64  `json:"speed"`
	MainOpen     bool     `json:"main_open"`
	BranchesOpen []string `json:"branches_open"`
}

type FiltrationReport struct {
	TankRuntime int     `json:"tank_runtime"`
	InFilter    bool    `json:"in_filter"`
	Pressure    float64 `json:"pressure"`
	MediaLife   int     `json:"media_life"`
}

type PumpDetail struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Running bool   `json:"running"`
	Cycles  int    `json:"cycles"`
}

type PumpsReport struct {
	Primary PumpDetail `json:"primary"`
	Standby PumpDetail `json:"standby"`
	Active  string     `json:"active"`
}

type ChemicalsReport struct {
	ChlorineDose float64            `json:"chlorine_dose"`
	Targets      []chlor.ZoneTarget `json:"targets"`
	PHTarget     float64            `json:"ph_target"`
	PHCurrent    float64            `json:"ph_current"`
	PHDirection  string             `json:"ph_direction"`
}

type TurbidityReport struct {
	Filtered  float64 `json:"filtered"`
	Raw       float64 `json:"raw"`
	SensorMax float64 `json:"sensor_max"`
}

type DrainReport struct {
	DrainOpen bool    `json:"drain_open"`
	Filling   bool    `json:"filling"`
	Target    float64 `json:"target"`
	Band      float64 `json:"band"`
}

type QuotaReport struct {
	Total       uint64             `json:"total"`
	Allocations []quota.Allocation `json:"allocations"`
}

type AuditReport struct {
	Total   int            `json:"total"`
	Volume  float64        `json:"volume"`
	Actions map[string]int `json:"actions"`
}

type AlarmItem struct {
	ID      string `json:"id"`
	Message string `json:"message"`
	Level   string `json:"level"`
}

type Snapshot struct {
	Pool        PoolReport        `json:"pool"`
	Circulation CirculationReport `json:"circulation"`
	Filtration  FiltrationReport  `json:"filtration"`
	Pumps       PumpsReport       `json:"pumps"`
	Chemicals   ChemicalsReport   `json:"chemicals"`
	Turbidity   TurbidityReport   `json:"turbidity"`
	Drain       DrainReport       `json:"drain"`
	Quota       QuotaReport       `json:"quota"`
	Audit       AuditReport       `json:"audit"`
	Alarms      []AlarmItem       `json:"alarms"`
}
