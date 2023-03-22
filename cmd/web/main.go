package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/k2realty/deals/pkg/config"
	"github.com/k2realty/deals/pkg/handlers"
	"github.com/k2realty/deals/pkg/render"
)

// portNumber is where we are serving our app
const portNumber = ":8080"

// app variable is our website configuration settings.
var app config.AppConfig
var session *scs.SessionManager

// main is the entry into our applictation
func main() {
	// ***SETTING*** set this to true when you move to production
	app.InProduction = false

	// below, we setup our sessions, and the appropriate cookie settings.
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	// below, we are storing that *scsSessionManager in our config variable.
	app.Session = session

	// creating the template cache so we do not have to read from disk each time someone sends a get request.
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	// storing the template cach in app config.
	app.TemplateCache = tc
	// ***SETTING***
	app.UseCache = false

	// these next two lines of code are sending our app config variable to handler package of our application.
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	// this next line of code sends our app config variable to the render package.
	render.NewTemplates(&app)

	fmt.Printf("Live server launched at http://localhost%v", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
