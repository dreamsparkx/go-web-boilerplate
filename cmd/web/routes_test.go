package main

import (
	"testing"

	"github.com/dreamsparkx/go-web-boilerplate/internal/config"
	gorilla "github.com/gorilla/mux"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig
	mux := routes(&app)
	switch v := mux.(type) {
	case *gorilla.Router:
		//do nothing
	default:
		t.Errorf("type is wrong, you are getting %T", v)
	}
}
