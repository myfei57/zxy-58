package console

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

type Pages struct {
	templates map[string]*template.Template
}

func NewPages(dir string) (*Pages, error) {
	p := &Pages{templates: map[string]*template.Template{}}
	for _, name := range []string{"pools", "quality", "chemicals", "alarms"} {
		path := filepath.Join(dir, name+".html")
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}
		t, err := template.New(name).Parse(string(data))
		if err != nil {
			return nil, err
		}
		p.templates[name] = t
	}
	return p, nil
}

func (p *Pages) Handle(name string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t, ok := p.templates[name]
		if !ok {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_ = t.Execute(w, map[string]string{"Page": name})
	}
}
