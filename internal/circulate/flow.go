package circulate

import "sync"

type Flow struct {
	mu      sync.RWMutex
	rate    float64
	running bool
	events  []string
}

func NewFlow(rate float64) *Flow {
	return &Flow{rate: rate, running: rate > 0}
}

func (f *Flow) Rate() float64 {
	f.mu.RLock()
	defer f.mu.RUnlock()
	return f.rate
}

func (f *Flow) SetRate(rate float64) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.rate = rate
}

func (f *Flow) Running() bool {
	f.mu.RLock()
	defer f.mu.RUnlock()
	return f.running
}

func (f *Flow) SetRunning(running bool) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.running == running {
		return
	}
	f.running = running
	if running {
		f.events = append(f.events, "flow-start")
	} else {
		f.events = append(f.events, "flow-stop")
	}
}

func (f *Flow) Events() []string {
	f.mu.RLock()
	defer f.mu.RUnlock()
	out := make([]string, len(f.events))
	copy(out, f.events)
	return out
}
