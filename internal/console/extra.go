package console

import "net/http"

func (s *Server) handleHealthDetail(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, s.Control.Health())
}

func (s *Server) handleHistory(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	writeJSON(w, s.Report.History(key))
}

func (s *Server) handleTrends(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, s.Report.Trends())
}

func (s *Server) handleAnalysis(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, s.Control.Analyze())
}

func (s *Server) handleAudit(w http.ResponseWriter, r *http.Request) {
	kind := r.URL.Query().Get("kind")
	writeJSON(w, map[string]any{
		"records": s.Control.Ledger.Query(kind),
		"volume":  s.Control.Ledger.TotalVolume(),
		"count":   s.Control.Ledger.ActionCount(kind),
	})
}
