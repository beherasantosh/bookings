package handlers

import (
	"net/http"

	"github.com/beherasantosh/bookings/pkg/config"
	"github.com/beherasantosh/bookings/pkg/models"
	"github.com/beherasantosh/bookings/pkg/render"
)

var Repo *Repository
type Repository struct{
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.SessionManager.Put(r.Context(), "remote_IP", remoteIP)

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}


func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello Santosh"


	remoteIP := m.App.SessionManager.GetString(r.Context(), "remote_IP")
	stringMap["remote_IP"] = remoteIP

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})

}