package console

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"poolops/internal/control"
	"poolops/internal/report"
)

type Server struct {
	Report  *report.Collector
	Control *control.Controller
	Pages   *Pages
}

func (s *Server) Router() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/health", s.handleHealth)
	r.Get("/pools", s.Pages.Handle("pools"))
	r.Get("/quality", s.Pages.Handle("quality"))
	r.Get("/chemicals", s.Pages.Handle("chemicals"))
	r.Get("/alarms", s.Pages.Handle("alarms"))
	r.Get("/api/report", s.handleReport)
	r.Get("/api/pools", s.handlePools)
	r.Get("/api/quality", s.handleQuality)
	r.Get("/api/chemicals", s.handleChemicals)
	r.Get("/api/alarms", s.handleAlarms)
	r.Get("/api/resolve", s.handleResolve)
	r.Get("/api/health", s.handleHealthDetail)
	r.Get("/api/history", s.handleHistory)
	r.Get("/api/trends", s.handleTrends)
	r.Get("/api/analysis", s.handleAnalysis)
	r.Get("/api/audit", s.handleAudit)
	r.Get("/api/stats", s.handleStats)
	r.Get("/api/zone", s.handleZone)
	r.Get("/api/inspect", s.handleInspect)
	r.Get("/api/store", s.handleStore)
	r.Get("/api/zones", s.handleZones)
	r.Get("/api/doses", s.handleDoses)
	r.Get("/api/pumps", s.handlePumps)
	r.Post("/api/cycle", s.handleCycle)
	r.Post("/api/pump", s.handlePump)
	r.Post("/api/treat-full", s.handleTreatFull)
	r.Post("/api/backwash", s.handleBackwash)
	r.Post("/api/failover", s.handleFailover)
	r.Post("/api/fill", s.handleFill)
	r.Post("/api/drain", s.handleDrain)
	r.Post("/api/treat", s.handleTreat)
	r.Post("/api/restore", s.handleRestore)
	return r
}
