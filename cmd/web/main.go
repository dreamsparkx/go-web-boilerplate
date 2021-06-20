package main

import (
	"net/http"
	"os"

	"github.com/alexedwards/scs/v2"
	"github.com/dreamsparkx/go-web-boilerplate/cmd/db"
	"github.com/dreamsparkx/go-web-boilerplate/cmd/env"
	webSession "github.com/dreamsparkx/go-web-boilerplate/cmd/web/session"
	"github.com/dreamsparkx/go-web-boilerplate/internal/config"
	"github.com/dreamsparkx/go-web-boilerplate/internal/handlers"
	"github.com/dreamsparkx/go-web-boilerplate/internal/render"
)

var session *scs.SessionManager
var app config.AppConfig

func main() {
	err := run()
	if err != nil {
		config.AppLogger.Fatalf("Error running application %s", err)
	}
	var mode string
	if len(os.Args) > 1 {
		mode = os.Args[1]
	}
	switch mode {
	case "api":
		modeAPI()
	case "migrate":
		err = db.MigrateDB(&app)
		if err != nil {
			config.AppLogger.Fatalf("Problem running migrations: %v", err)
		}
	default:
		config.AppLogger.Fatalf("Invalid mode %s, must be either 'api' or 'migrate'")
	}
}

func modeAPI() {
	srv := &http.Server{
		Addr:    ":" + app.Port,
		Handler: routes(&app),
	}
	config.AppLogger.Infof("Starting Application on port: %s", app.Port)
	err := srv.ListenAndServe()
	config.AppLogger.Fatalf("Server Creation Failed: %s", err)
}

func run() error {
	env.LoadENV(&app)
	session = webSession.CreateSession(&app)
	app.Session = session
	config.InitLogger(app.InProduction)
	defer config.AppLogger.Sync()
	db.ConnectDB(&app)
	tc, err := render.CreateTemplateCache()
	if err != nil {
		config.AppLogger.Fatal("Cannot Create Template Cache")
		return err
	}
	app.TemplateCache = tc
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)
	return err
}

// https://github.com/cogolabs/go-boilerplate
// https://github.com/vardius/go-api-boilerplate
