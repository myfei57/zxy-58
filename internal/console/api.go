package console

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func writeJSON(w http.ResponseWriter, value any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_ = json.NewEncoder(w).Encode(value)
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, map[string]string{"status": "ok"})
}

func (s *Server) handleReport(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, s.Report.Collect())
}

func (s *Server) handlePools(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, s.Report.Collect().Pool)
}

func (s *Server) handleQuality(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, s.Report.Quality())
}

func (s *Server) handleChemicals(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, s.Report.Collect().Chemicals)
}

func (s *Server) handleAlarms(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, s.Report.Collect().Alarms)
}

func (s *Server) handleResolve(w http.ResponseWriter, r *http.Request) {
	probeID := r.URL.Query().Get("probe")
	writeJSON(w, map[string]string{"zone": s.Control.ResolveZone(probeID)})
}

func (s *Server) handleBackwash(w http.ResponseWriter, r *http.Request) {
	if err := s.Control.Backwash(); err != nil {
		writeJSON(w, map[string]any{"error": err.Error()})
		return
	}
	writeJSON(w, map[string]string{"status": "ok"})
}

func (s *Server) handleFailover(w http.ResponseWriter, r *http.Request) {
	if err := s.Control.Failover(); err != nil {
		writeJSON(w, map[string]any{"error": err.Error()})
		return
	}
	writeJSON(w, map[string]string{"status": "ok"})
}

func (s *Server) handleFill(w http.ResponseWriter, r *http.Request) {
	zoneID := r.URL.Query().Get("zone")
	amount, err := s.Control.Fill(zoneID)
	if err != nil {
		writeJSON(w, map[string]any{"error": err.Error()})
		return
	}
	writeJSON(w, map[string]any{"amount": amount})
}

func (s *Server) handleDrain(w http.ResponseWriter, r *http.Request) {
	zoneID := r.URL.Query().Get("zone")
	level, _ := strconv.ParseFloat(r.URL.Query().Get("level"), 64)
	if err := s.Control.Drain(zoneID, level); err != nil {
		writeJSON(w, map[string]any{"error": err.Error()})
		return
	}
	writeJSON(w, map[string]string{"status": "ok"})
}

func (s *Server) handleTreat(w http.ResponseWriter, r *http.Request) {
	hour, _ := strconv.Atoi(r.URL.Query().Get("hour"))
	raw, _ := strconv.ParseFloat(r.URL.Query().Get("raw"), 64)
	if err := s.Control.Treat(hour, raw); err != nil {
		writeJSON(w, map[string]any{"error": err.Error()})
		return
	}
	writeJSON(w, map[string]string{"status": "ok"})
}

func (s *Server) handleRestore(w http.ResponseWriter, r *http.Request) {
	branches := r.URL.Query()["zone"]
	if err := s.Control.Restore(branches); err != nil {
		writeJSON(w, map[string]any{"error": err.Error()})
		return
	}
	writeJSON(w, map[string]string{"status": "ok"})
}

func (s *Server) handleCycle(w http.ResponseWriter, r *http.Request) {
	hour, _ := strconv.Atoi(r.URL.Query().Get("hour"))
	writeJSON(w, s.Control.RunCycle(hour))
}
