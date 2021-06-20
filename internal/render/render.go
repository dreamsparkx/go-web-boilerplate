package render

import (
	"bytes"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/dreamsparkx/go-web-boilerplate/internal/config"
	"github.com/dreamsparkx/go-web-boilerplate/internal/models"
	"github.com/justinas/nosurf"
)

var functions = template.FuncMap{}

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.CSRFToken = nosurf.Token(r)
	return td
}

func RenderTemplate(rw http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {
	var tc config.TemplateCache
	if app.UseTemplateCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}
	t, ok := tc[tmpl]
	if !ok {
		config.AppLogger.Fatal("Could not get template from TemplateCache")
	}
	buf := new(bytes.Buffer)
	td = AddDefaultData(td, r)
	_ = t.Execute(buf, td)
	_, err := buf.WriteTo(rw)
	if err != nil {
		config.AppLogger.Errorf("error writing template to browser: %s", err)
	}
}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (config.TemplateCache, error) {
	myCache := config.TemplateCache{}
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		config.AppLogger.Infof("Creating Template %s", page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page) //create template set
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/layouts/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/layouts/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
