package update

import (
	"github.com/Krab1o/meebin/internal/model/event/dto"
	smodel "github.com/Krab1o/meebin/internal/model/event/s_model"
)

func UpdatedEventDTOToService(event *dto.UpdateEvent) *smodel.Event {
	data := UpdatedEventDataDTOToService(event.Data)
	return &smodel.Event{
		Id:     event.Id,
		Status: event.Status,
		Data:   data,
	}
}

func UpdatedEventDataDTOToService(data *dto.UpdateEventData) *smodel.EventData {
	return &smodel.EventData{
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
