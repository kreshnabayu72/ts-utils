package service

import (
	"errors"
	"fmt"
)

type RegisteredService struct {
	Name    string
	Service any
}

type ServiceManager struct {
	RegisteredServices map[string]RegisteredService
}

func (s *ServiceManager) Register(name string, service any) {
	s.RegisteredServices[name] = RegisteredService{
		Name:    name,
		Service: service,
	}
}

func (s *ServiceManager) Unregister(name string) {
	delete(s.RegisteredServices, name)
}

func (s *ServiceManager) GetService(name string) (any, error) {
	if s.RegisteredServices == nil {
		return nil, errors.New("no service has been registered yet")
	}

	r, ok := s.RegisteredServices[name]
	if !ok {
		return nil, errors.New(fmt.Sprintf("service '%s' is not registered", name))
	}

	return r.Service, nil
}
