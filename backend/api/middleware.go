package api

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

// setupLoggingMiddleware is a function that initializes the loggin middleware
func (s *Server) setupLoggingMiddleware() {
	// add logger middleware
	s.router.Use(func(h http.Handler) http.Handler {
		return handlers.LoggingHandler(os.Stdout, h)
	})
}

// recoveryMiddleware is a function that initializes the recovery middleware
func (s *Server) setupRecoveryMiddleware() {
	// add logger middleware
	s.router.Use(handlers.RecoveryHandler(handlers.PrintRecoveryStack(true)))
}
