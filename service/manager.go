package service

import (
	"fmt"
	"log/slog"
)

type Servicer interface {
	// Init executes right after service is registered to service manager. Servicer should panic if encounters any error.
	Init()
	// Close executes right after service is unregistered from service manager. Servicer should NOT panic if encounters any error.
	Close()
}

// RegisteredService is the registered service.
type RegisteredService struct {
	Name    string
	Service Servicer
}

// Managerer is the interface for implementing inject/provide management system.
type Managerer interface {
	Register(name string, service Servicer)
	Unregister(name string)
	Get(name string) Servicer
}

// Manager is the manager for registration, unregistration and retrieval of services.
type Manager struct {
	RegisteredServices map[string]RegisteredService
}

// MakeManager creates service manager instance.
func MakeManager() Managerer {
	return &Manager{
		RegisteredServices: make(map[string]RegisteredService),
	}
}

// Register registers service.
//
// if registered service already exists with the same name, it will be overridden with the new one.
func (s *Manager) Register(name string, service Servicer) {
	s.RegisteredServices[name] = RegisteredService{
		Name:    name,
		Service: service,
	}

	rs := s.RegisteredServices[name]

	slog.Debug(fmt.Sprintf("Initializing service '%s'...", rs.Name))
	rs.Service.Init()
	slog.Debug(fmt.Sprintf("Successfully initialized service '%s'.", rs.Name))
}

// Unregister unregisters registered service.
//
// This function can be safely called more than once, even if the service has not been registered yet.
func (s *Manager) Unregister(name string) {
	rs, ok := s.RegisteredServices[name]
	if !ok {
		// Pass
		return
	}

	slog.Debug(fmt.Sprintf("Closing service '%s'...", rs.Name))
	rs.Service.Close()
	slog.Debug(fmt.Sprintf("Successfully closed service '%s'.", rs.Name))
	delete(s.RegisteredServices, name)
}

// Get retrieves the registered service.
//
// May panic if service has not been registered yet.
func (s *Manager) Get(name string) Servicer {
	if s.RegisteredServices == nil {
		panic("no service has been registered yet")
	}

	r, ok := s.RegisteredServices[name]
	if !ok {
		panic(fmt.Sprintf("service '%s' is not registered", name))
	}

	return r.Service
}
