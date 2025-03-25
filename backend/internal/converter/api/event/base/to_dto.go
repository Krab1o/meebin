package base

import (
	"github.com/Krab1o/meebin/internal/model/event/dto"
	smodel "github.com/Krab1o/meebin/internal/model/event/s_model"
)

func EventServiceToDTO(event *smodel.Event) *dto.BaseEvent {
	data := EventDataServiceToDTO(event.Data)
	return &dto.BaseEvent{
		Id:     event.Id,
		Status: event.Status,
		Data:   data,
	}
}

func EventDataServiceToDTO(data *smodel.EventData) *dto.BaseEventData {
	return &dto.BaseEventData{
		Latitude:     data.Latitude,
		Longtitude:   data.Longtitude,
		Title:        data.Title,
		Description:  data.Description,
		CallerId:     data.CallerId,
		TimeCalled:   data.TimeCalled,
		UtilizatorId: data.UtilizatorId,
		TimeUtilized: data.TimeUtilized,
	}
}
