package event

import (
	rmodel "github.com/Krab1o/meebin/internal/model/event/r_model"
	smodel "github.com/Krab1o/meebin/internal/model/event/s_model"
)

func EventServiceToRepo(event *smodel.Event) *rmodel.Event {
	repoData := EventDataServiceToRepo(event.Data)
	return &rmodel.Event{
		Id:     event.Id,
		Status: event.Status,
		Data:   repoData,
	}
}

func EventDataServiceToRepo(data *smodel.EventData) *rmodel.EventData {
	if data == nil {
		return nil
	}
	return &rmodel.EventData{
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
