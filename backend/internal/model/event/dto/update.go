package dto

import (
	"time"

	"github.com/Krab1o/meebin/internal/model"
)

type UpdateEvent struct {
	Id     uint64            `json:"id,omitempty" example:"42"`
	Status model.EventStatus `json:"status"`
	Data   *UpdateEventData  `json:"data"`
}

type UpdateEventData struct {
	Latitude     float64   `json:"latitude"      example:"42.881"`
	Longtitude   float64   `json:"longtitude"    example:"45.254"`
	Title        string    `json:"title"         example:"Ужасное загрязнение"`
	Description  string    `json:"description"   example:"Бла-бла-бла"`
	CallerId     uint64    `json:"caller_id"     example:"42"`
	TimeCalled   time.Time `json:"time_called"                                 format:"date-time"`
	UtilizatorId uint64    `json:"utilizator_id" example:"42"`
	TimeUtilized time.Time `json:"time_utilized"                               format:"date-time"`
}
