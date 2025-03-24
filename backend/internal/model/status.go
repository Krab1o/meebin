package model

type EventStatus uint64

// Correlates with event_status_ id from sql table
const (
	StatusOnModeration EventStatus = iota + 1
	StatusOpened
	StatusOnConfiramation
	StatusClosed
)
