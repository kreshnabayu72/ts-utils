package service

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"gitea.qwertysystem.net/BETS/go-agenda"
	"gitea.qwertysystem.net/BETS/go-agenda/mongo"
)

// AgendaService abstracts the agenda.Agenda implementation as a service
type AgendaService interface {
	Start()
	Define(name string, jobFunc func(map[string]any) error)
	Schedule(name string, t time.Time, data map[string]any)
	Stop(filter bson.M)
}

// AgendaServiceImpl is the implementation of the AgendaService
type AgendaServiceImpl struct {
	CollectionName string
	Agenda         *agenda.Agenda
}

// MakeAgendaService creates AgendaServiceImpl
func MakeAgendaService(collectionName string) *AgendaServiceImpl {
	var err error

	db := mongo.MakeCollection(collectionName)

	a, err := agenda.MakeAgenda(db)
	if err != nil {
		panic(err)
	}

	return &AgendaServiceImpl{
		CollectionName: collectionName,
		Agenda:         a,
	}
}

// Define defines the job that will be watched by the agenda
//
// Mainly called in initialization.
// You only need to call this once per job, or when you wanted to extend jobFunc with another function call
func (s *AgendaServiceImpl) Define(name string, jobFunc func(map[string]any) error) {
	if err := s.Agenda.Define(name, jobFunc); err != nil {
		panic(err)
	}
}

// Schedule defines a schedule of the job, with provided name and time, along with optional data
func (s *AgendaServiceImpl) Schedule(name string, t time.Time, data map[string]any) {
	if err := s.Agenda.NewSchedule(name, t, data); err != nil {
		panic(err)
	}
}

// Stop stops any running schedule with the provided filter
func (s *AgendaServiceImpl) Stop(filter bson.M) {
	if err := s.Agenda.Stop(filter); err != nil {
		panic(err)
	}
}

// Start starts running and watching scheduled job(s)
func (s *AgendaServiceImpl) Start() {
	if err := s.Agenda.Start(); err != nil {
		panic(err)
	}
}
