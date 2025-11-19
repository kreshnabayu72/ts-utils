package provider

import (
	"github.com/stretchr/testify/mock"

	"gitea.qwertysystem.net/BETS/ts-utils/v2/service"
)

type AgendaServiceProvider interface {
	MakeAgendaService(collectionName string) service.AgendaService
}

type AgendaServiceProviderImpl struct{}

func (asf *AgendaServiceProviderImpl) MakeAgendaService(collectionName string) service.AgendaService {
	return service.MakeAgendaService(collectionName)
}

type MockAgendaServiceProvider struct {
	mock.Mock
}

func (m *MockAgendaServiceProvider) MakeAgendaService(collectionName string) service.AgendaService {
	args := m.Called(collectionName)
	return args.Get(0).(service.AgendaService)
}
