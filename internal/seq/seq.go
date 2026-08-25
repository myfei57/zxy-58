package seq

type Recorder struct {
	Events []string
}

func NewRecorder() *Recorder {
	return &Recorder{}
}

func (r *Recorder) Add(event string) {
	r.Events = append(r.Events, event)
}

func (r *Recorder) Last() string {
	if len(r.Events) == 0 {
		return ""
	}
	return r.Events[len(r.Events)-1]
}
