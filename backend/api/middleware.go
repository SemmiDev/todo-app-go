package api

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

// InitMiddlewares is a function that initializes the middlewares
func (s *Server) setupMiddleware() {
	s.loggingMiddleware()
	s.recoveryMiddleware()
}

// initMiddlewares is a function that initializes the middlewares
func (s *Server) loggingMiddleware() {
	// add logger middleware
	s.router.Use(func(h http.Handler) http.Handler {
		return handlers.LoggingHandler(os.Stdout, h)
	})
}

// recoveryMiddleware is a function that initializes the recovery middleware
func (s *Server) recoveryMiddleware() {
	// add logger middleware
	s.router.Use(handlers.RecoveryHandler(handlers.PrintRecoveryStack(true)))
}
