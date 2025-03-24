package model

type EventStatus uint64

// Correlates with event_status_ id from sql table
const (
	StatusOnModeration    EventStatus = 1
	StatusOpened                      = 2
	StatusOnConfiramation             = 3
	StatusClosed                      = 4
)
