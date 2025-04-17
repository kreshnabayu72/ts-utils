package service

import (
	"time"

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

func (s *AgendaServiceImpl) Schedule(name string, t time.Time, data map[string]any) {

}
