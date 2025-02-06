package event

type EventAPI interface {
	Get()
	GetById()
	Create()
	Update()
	Delete()
}