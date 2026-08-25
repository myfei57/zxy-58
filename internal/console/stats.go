package console

import "net/http"

func (s *Server) handleStats(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, map[string]any{
		"pool":   s.Report.PoolStats(),
		"flow":   s.Report.FlowStats(),
		"health": s.Report.HealthSummary(),
	})
}

func (s *Server) handleZone(w http.ResponseWriter, r *http.Request) {
	zoneID := r.URL.Query().Get("zone")
	writeJSON(w, map[string]any{
		"dose":     s.Control.ZoneDose(zoneID),
		"fill_due": s.Control.ZoneFillDue(zoneID),
	})
}

func (s *Server) handleInspect(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, s.Control.Inspect())
}
