package pump

type Telemetry struct {
	RuntimeHours int
	StartCount   int
	StopCount    int
}

func (t *Telemetry) RecordStart() {
	t.StartCount++
}

func (t *Telemetry) RecordStop() {
	t.StopCount++
}

func (t *Telemetry) AddRuntime(hours int) {
	t.RuntimeHours += hours
}

func (t Telemetry) Cycles() int {
	if t.StartCount < t.StopCount {
		return t.StartCount
	}
	return t.StopCount
}
