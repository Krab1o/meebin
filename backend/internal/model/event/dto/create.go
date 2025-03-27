package dto

type NewEvent struct {
	Id   uint64        `json:"id,omitempty" example:"42" binding:"required"`
	Data *NewEventData `json:"data"                      binding:"required"`
}

type NewEventData struct {
	Latitude    float64 `json:"latitude"    example:"42.881"              binding:"required"`
	Longtitude  float64 `json:"longtitude"  example:"45.254"              binding:"required"`
	Title       string  `json:"title"       example:"Ужасное загрязнение" binding:"required"`
	Description string  `json:"description" example:"Бла-бла-бла"         binding:"required"`
	// CallerId    uint64    `json:"caller_id"   example:"42"                  binding:"required"`
	// TimeCalled  time.Time `json:"time_called"                               binding:"required" format:"date-time"`
	// UtilizatorId uint64    `json:"utilizator_id" example:"42"`
	// TimeUtilized time.Time `json:"time_utilized"                                                  format:"date-time"`
}
