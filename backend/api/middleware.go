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

// setupCORS Optional, CORS config, to make sure it can be called from everywhere
func (s *Server) setupCORS() {
	headersOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	s.router.Use(handlers.CORS(headersOk, methodsOk))
}
