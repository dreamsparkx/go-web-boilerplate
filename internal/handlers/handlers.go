package handlers

import (
	"net/http"

	"github.com/dreamsparkx/go-web-boilerplate/internal/config"
	"github.com/dreamsparkx/go-web-boilerplate/internal/models"
	"github.com/dreamsparkx/go-web-boilerplate/internal/render"
)

type repository struct {
	App *config.AppConfig
}

// Repo the repository used by handlers
var Repo *repository

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *repository {
	return &repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *repository) {
	Repo = r
}

func (repo *repository) Home(rw http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	repo.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(rw, r, "home.page.tmpl", &models.TemplateData{})
}

func (repo *repository) About(rw http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"
	stringMap["remote_ip"] = repo.App.Session.GetString(r.Context(), "remote_ip")
	render.RenderTemplate(rw, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
