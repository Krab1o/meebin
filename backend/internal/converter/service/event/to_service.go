package event

import (
	rmodel "github.com/Krab1o/meebin/internal/model/event/r_model"
	smodel "github.com/Krab1o/meebin/internal/model/event/s_model"
)

func EventRepoToService(event *rmodel.Event) *smodel.Event {
	serviceData := EventDataRepoToService(event.Data)
	return &smodel.Event{
		Id:     event.Id,
		Status: event.Status,
		Data:   serviceData,
	}
}

func EventDataRepoToService(data *rmodel.EventData) *smodel.EventData {
	if data == nil {
		return nil
	}
	return &smodel.EventData{
		Latitude:     data.Latitude,
		Longtitude:   data.Longtitude,
		Title:        data.Title,
		Description:  data.Description,
		CallerId:     data.CallerId,
		UtilizatorId: data.UtilizatorId,
		TimeCalled:   data.TimeCalled,
		TimeUtilized: data.TimeUtilized,
	}
}
