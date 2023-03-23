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

// NewRepo our app config variable, and stores it in a repo.
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers assigns our Local Repo var to the incoming pointer to r.
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remoteIP", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the Deals page handler
func (m *Repository) Deals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "deals.page.tmpl", &models.TemplateData{})
}

// About is the PostContact page handler
func (m *Repository) PostDeals(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Posted to contact"))
}

// Home is the home page handler
func (m *Repository) Login(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "login.page.tmpl", &models.TemplateData{})
}
