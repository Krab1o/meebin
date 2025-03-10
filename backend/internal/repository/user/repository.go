package user

import (
	rmodel "github.com/Krab1o/meebin/internal/model/r_model"
	"github.com/Krab1o/meebin/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) repository.UserRepository {
	return &repo{db: db}
}

// TODO: implement repo get one user logic
func (r *repo) GetById(id int64) (*rmodel.User, error) {
	return nil, nil
}

// TODO: implement repo get users logic
func (r *repo) List() ([]rmodel.User, error) {
	return nil, nil
}

// TODO: implement repo update user logic
func (r *repo) Update(id int64) error {
	return nil
}

// TODO: implement repo delete user logic
func (r *repo) Delete(id int64) error {
	return nil
}
