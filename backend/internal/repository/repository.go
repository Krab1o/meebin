package repository

import (
	event "github.com/Krab1o/meebin/internal/repository/event/model"
	user "github.com/Krab1o/meebin/internal/repository/user/model"
)

type UserRepository interface {
	Add(*user.User) (int64, error)
	GetById(id int64) (*user.User, error)
	List() ([]user.User, error)
	Update(id int64) error
	Delete(id int64) error
}

type EventRepository interface {
	Add(*event.Event) (int64, error)
	GetById(id int64) (*event.Event, error)
	List() ([]event.Event, error)
	Update(id int64) error
	Delete(id int64) error
}
