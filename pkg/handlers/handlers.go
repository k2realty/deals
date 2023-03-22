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
	// perform some logic
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remoteIP", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the contact page handler
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "The lucky number 11"

	remoteIP := m.App.Session.GetString(r.Context(), "remoteIP")
	stringMap["remoteIP"] = remoteIP

	render.RenderTemplate(w, "contact.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Home is the home page handler
func (m *Repository) Login(w http.ResponseWriter, r *http.Request) {
	// perform some logic

	render.RenderTemplate(w, "login.page.tmpl", &models.TemplateData{})
}
