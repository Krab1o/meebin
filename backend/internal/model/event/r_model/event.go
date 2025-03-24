package rmodel

import (
	"time"

	"github.com/Krab1o/meebin/internal/model"
)

type Event struct {
	Id         uint64
	CallerId   uint64
	UtilizerId uint64
	Status     model.EventStatus
	Data       *EventData
}

type EventData struct {
	Latitude    float64
	Longtitude  float64
	Title       string
	Description string
	TimeCalled  time.Time
	TimeCleaned time.Time
}
