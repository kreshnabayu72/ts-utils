package service

import (
	"time"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
)

type MockAgendaService struct {
	mock.Mock
}

func (m *MockAgendaService) Start() {
	m.Called()
}

func (m *MockAgendaService) Define(name string, jobFunc func(map[string]any) error) {
	m.Called(name, jobFunc)
}

func (m *MockAgendaService) Schedule(name string, t time.Time, data map[string]any) {
	m.Called(name, t, data)
}

func (m *MockAgendaService) Every(name string, t time.Time, data map[string]any, repeatEvery string) {
	m.Called(name, t, data, repeatEvery)
}

func (m *MockAgendaService) Stop(filter bson.M) {
	m.Called(filter)
}
