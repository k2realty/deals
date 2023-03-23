package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/k2realty/deals/internal/config"
	"github.com/k2realty/deals/internal/models"
	"github.com/k2realty/deals/internal/render"
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

	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

// About is the Deals page handler
func (m *Repository) Deals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "deals.page.tmpl", &models.TemplateData{})
}

// About is the PostDeals page handler
func (m *Repository) PostDeals(w http.ResponseWriter, r *http.Request) {
	// when you pull data out of a form request, it's always a string.
	clientName := r.Form.Get("clientName")
	comp := r.Form.Get("comp")

	w.Write([]byte(fmt.Sprintf("The client: %s, has added $%v to your pipeline!", clientName, comp)))
}

type jsonResponse struct {
	ClientFName string `JSON:"clientFName"`
	ClientLName string `JSON:"clientLName"`
}

// About is the DealsJSON page handler
func (m *Repository) DealsJSON(w http.ResponseWriter, r *http.Request) {

	resp := jsonResponse{
		ClientFName: "Dan",
		ClientLName: "Bogart",
	}

	out, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Println(err)
	}
	// this tells the browser what kind of information we are sending.
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// Home is the home page handler
func (m *Repository) Login(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "login.page.tmpl", &models.TemplateData{})
}
