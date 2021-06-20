package config

import (
	"html/template"

	"github.com/alexedwards/scs/v2"
)

type TemplateCache map[string]*template.Template

type AppConfig struct {
	TemplateCache    TemplateCache
	UseTemplateCache bool
	InProduction     bool
	Session          *scs.SessionManager
	Port             string
	DB               *Database
	DbCredentials    struct {
		Host            string
		Name            string
		User            string
		Password        string
		SSLModeDisabled string
	}
}
