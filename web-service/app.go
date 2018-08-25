package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// App struct
type App struct {
	Address string
	Router  *mux.Router
	Server  *http.Server
}

// NewApp constructor
func NewApp(
	host string,
	port int,
) (*App, error) {
	a := &App{
		Address: fmt.Sprintf("%s:%d", host, port),
	}
	a.configureApp()
	return a, nil
}

func (a *App) configureApp() {
	a.Router = a.getRouter()
	a.configureServer()
}

func (a *App) getRouter() *mux.Router {
	router := mux.NewRouter()
	router.Handle("/api/get", NewServersHandler(a)).Methods("GET")
	router.Handle("/api/post", NewPushToUsersHandler(a)).Methods("POST")
	return router
}

func (a *App) configureServer() {
	a.Server = &http.Server{
		Addr:    a.Address,
		Handler: a.Router,
	}
}

// Init starts the app
func (a *App) Init() {
	go a.Server.ListenAndServe()
}
