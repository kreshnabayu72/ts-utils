package service

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"gitea.qwertysystem.net/BETS/go-agenda"
	"gitea.qwertysystem.net/BETS/go-agenda/mongo"
)

type AgendaService interface{}

type AgendaServiceImpl struct {
	CollectionName string
	Agenda         *agenda.Agenda
}

func MakeAgendaService(collectionName string) *AgendaServiceImpl {
	var err error

	db := mongo.MakeCollection(collectionName)

	a, err := agenda.New(db)
	if err != nil {
		panic(err)
	}

	return &AgendaServiceImpl{
		CollectionName: collectionName,
		Agenda:         a,
	}
}

func (s *AgendaServiceImpl) Define(name string, jobFunc func(map[string]any) error) {
	if err := s.Agenda.Define(name, jobFunc); err != nil {
		panic(err)
	}
}

func (s *AgendaServiceImpl) Schedule(name string, t time.Time, data map[string]any) {
	if err := s.Agenda.NewSchedule(name, t, data); err != nil {
		panic(err)
	}
}

func (s *AgendaServiceImpl) Stop(filter bson.M) {
	if err := s.Agenda.Stop(filter); err != nil {
		panic(err)
	}
}
