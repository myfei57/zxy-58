package console

import (
	"net/http"
	"strconv"
)

func (s *Server) handleStore(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, s.Report.Store.Snapshot())
}

func (s *Server) handleZones(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, s.Report.ZoneDetails())
}

func (s *Server) handleDoses(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, map[string]any{
		"doses": s.Control.PerZoneDose(),
		"total": s.Control.TotalDose(),
	})
}

func (s *Server) handlePumps(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, s.Control.PumpStatuses())
}

func (s *Server) handlePump(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	action := r.URL.Query().Get("action")
	ok := false
	if action == "start" {
		ok = s.Control.StartPump(id)
	} else if action == "stop" {
		ok = s.Control.StopPump(id)
	}
	writeJSON(w, map[string]any{"id": id, "action": action, "ok": ok})
}

func (s *Server) handleTreatFull(w http.ResponseWriter, r *http.Request) {
	hour, _ := strconv.Atoi(r.URL.Query().Get("hour"))
	writeJSON(w, s.Control.FullTreat(hour))
}
