package create

import (
	"github.com/Krab1o/meebin/internal/model/event/dto"
	smodel "github.com/Krab1o/meebin/internal/model/event/s_model"
)

func NewEventServiceToDTO(event *dto.NewEvent) *smodel.Event {
	data := NewEventDataServiceToDTO(event.Data)
	return &smodel.Event{
		Id:   event.Id,
		Data: data,
	}
}

func NewEventDataServiceToDTO(event *dto.NewEventData) *smodel.EventData {
	return &smodel.EventData{
		Latitude:    event.Latitude,
		Longtitude:  event.Longtitude,
		Title:       event.Title,
		Description: event.Description,
		// CallerId:    event.CallerId,
		// TimeCalled:  event.TimeCalled,
	}
}
