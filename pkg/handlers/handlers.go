package handlers

import (
	"net/http"

	"github.com/k2realty/deals/pkg/config"
	"github.com/k2realty/deals/pkg/models"
	"github.com/k2realty/deals/pkg/render"
)

// repository pattern. Allows swapping components with little effort
var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remoteIP", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "The lucky number 11"

	remoteIP := m.App.Session.GetString(r.Context(), "remoteIP")
	stringMap["remoteIP"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
