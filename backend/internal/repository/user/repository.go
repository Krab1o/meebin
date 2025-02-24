package user

import (
	"github.com/Krab1o/meebin/internal/repository"
	repoModel "github.com/Krab1o/meebin/internal/repository/user/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

//TODO: add constants for table names and columns

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) repository.UserRepository {
	return &repo{db: db}
}

// TODO: implement repo add user logic
func (r *repo) Add(*repoModel.User) (int64, error) {
	return 0, nil
}

// TODO: implement repo get one user logic
func (r *repo) GetById(id int64) (*repoModel.User, error) {
	return nil, nil
}

// TODO: implement repo get users logic
func (r *repo) List() ([]repoModel.User, error) {
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
