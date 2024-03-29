package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/beherasantosh/bookings/pkg/config"
	"github.com/beherasantosh/bookings/pkg/handlers"
	"github.com/beherasantosh/bookings/pkg/render"
)

const portNumber = ":8080"
var app config.AppConfig
var sessionManager *scs.SessionManager


func main() {
	app.InProduction = false

	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	sessionManager.Cookie.Secure = app.InProduction

	app.SessionManager = sessionManager

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)
	

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))

	srv := &http.Server{
		Addr:              portNumber,
		Handler:           routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
