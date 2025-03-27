package smodel

import (
	"time"

	"github.com/Krab1o/meebin/internal/model"
)

type Event struct {
	Id     uint64
	Status model.EventStatus
	Data   *EventData
}

type EventData struct {
	Latitude     float64
	Longtitude   float64
	Title        string
	Description  string
	CallerId     uint64
	UtilizatorId uint64
	TimeCalled   time.Time
	TimeUtilized time.Time
}
