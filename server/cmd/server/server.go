package main

import (
	"context"
	"fmt"
	"github.com/rebase/architecture-lab2/server/balancers"
	"net/http"
)

type HttpPortNumber int

// BalanceApiServer configures necessary handlers and starts listening on a configured port.
type BalanceApiServer struct {
	Port HttpPortNumber

	BalancersHandler balancers.HttpHandlerFunc

	server *http.Server
}

// Start will set all handlers and start listening.
// If this methods succeeds, it does not return until server is shut down.
// Returned error will never be nil.
func (s *BalanceApiServer) Start() error {
	if s.BalancersHandler == nil {
		return fmt.Errorf("balancers HTTP handler is not defined - cannot start")
	}
	if s.Port == 0 {
		return fmt.Errorf("port is not defined")
	}

	handler := new(http.ServeMux)
	handler.HandleFunc("/balancers", s.BalancersHandler)

	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.Port),
		Handler: handler,
	}

	return s.server.ListenAndServe()
}

// Stops will shut down previously started HTTP server.
func (s *BalanceApiServer) Stop() error {
	if s.server == nil {
		return fmt.Errorf("server was not started")
	}
	return s.server.Shutdown(context.Background())
}
