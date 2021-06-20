package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dreamsparkx/go-web-boilerplate/internal/config"
	"github.com/dreamsparkx/go-web-boilerplate/internal/handlers"
	muxHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func routes(app *config.AppConfig) http.Handler {
	router := mux.NewRouter()
	router.Use(muxHandlers.RecoveryHandler())
	router.Use(muxHandlers.CompressHandler)
	router.Use(NoSurf)
	router.Use(SessionLoad)
	router.HandleFunc("/", handlers.Repo.Home).Methods("GET")
	router.HandleFunc("/about", handlers.Repo.About).Methods("GET")
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		notFoundHandler(w, router)
	})
	addStaticFiles(router)
	return router
}

func addStaticFiles(router *mux.Router) {
	s := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	router.PathPrefix("/static/").Handler(s)
	http.Handle("/", router)
}

func notFoundHandler(w http.ResponseWriter, router *mux.Router) {
	fmt.Fprintf(w, "Available Routes:\n")
	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		methods, _ := route.GetMethods()
		path, _ := route.GetPathTemplate()
		queries, _ := route.GetQueriesTemplates()
		queryParams := ""
		if len(queries) > 0 {
			queryParams = fmt.Sprintf("?%s", strings.Join(queries, ","))
		}
		fmt.Fprintf(w, "%s %s%s\n", methods, path, queryParams)
		return nil
	})
}

// csrf_token=NlZ3ZZImaSnbhTmyt4SqnDv7nUlvXU8Wc8Xlv4Nhm5s=; session=o5QFB6tgLRTGm6gtqSHncoBI-oyw22N1ZF8P1a2oKHE
