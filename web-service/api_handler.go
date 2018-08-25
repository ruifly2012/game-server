package main

import (
	"net/http"
)

// ServersHandler handler
type ServersHandler struct {
	App *App
}

// NewServersHandler creates a new servers handler
func NewServersHandler(a *App) *ServersHandler {
	m := &ServersHandler{App: a}
	return m
}

// ServeHTTP method
func (s *ServersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello"))
}
