package service

import (
	"fmt"
)

// RegisteredService is the registered service.
//
// TODO(refan): Might be better off as an interface, so every service must comply to the interface's implementation.
type RegisteredService struct {
	Name    string
	Service any
}

// ServiceManager is the manager for registration, unregistration and retrieval of services.
type ServiceManager struct {
	RegisteredServices map[string]RegisteredService
}

// Register registers service.
//
// if registered service already exists with the same name, it will be overridden with the new one.
func (s *ServiceManager) Register(name string, service any) {
	s.RegisteredServices[name] = RegisteredService{
		Name:    name,
		Service: service,
	}
}

// Unregister unregisters registered service.
// This function can be safely called even if the service has not registered yet.
func (s *ServiceManager) Unregister(name string) {
	delete(s.RegisteredServices, name)
}

// GetService retrieves the registered service.
//
// May panic if service has not been registered yet.
func (s *ServiceManager) GetService(name string) any {
	if s.RegisteredServices == nil {
		panic("no service has been registered yet")
	}

	r, ok := s.RegisteredServices[name]
	if !ok {
		panic(fmt.Sprintf("service '%s' is not registered", name))
	}

	return r.Service
}
