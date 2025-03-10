package role

import (
	"context"
	"fmt"

	"github.com/Krab1o/meebin/internal/model"
	"github.com/Krab1o/meebin/internal/repository"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func (r *repo) GetUserRolesById(
	ctx context.Context,
	tx pgx.Tx,
	userId uint64,
) ([]model.Role, error) {
	query, args, err := squirrel.Select(
		repository.RoleTitleColumn,
	).
		PlaceholderFormat(squirrel.Dollar).
		From(repository.UserRoleTableName).
		LeftJoin(fmt.Sprintf("%s ON %s.%s = %s.%s",
			repository.RoleTableName,
			repository.RoleTableName,
			repository.RoleIdColumn,
			repository.UserRoleTableName,
			repository.UserRoleIdColumn,
		),
		).
		ToSql()
	if err != nil {
		return nil, repository.NewInternalError(err)
	}
	var rows pgx.Rows
	if tx != nil {
		rows, err = tx.Query(ctx, query, args...)
	} else {
		rows, err = r.db.Query(ctx, query, args...)
	}
	if err != nil {
		return nil, repository.NewInternalError(err)
	}

	var roles []model.Role
	var role model.Role
	for rows.Next() {
		err = rows.Scan(&role)
		if err != nil {
			return nil, repository.NewInternalError(err)
		}
		roles = append(roles, role)
	}

	return roles, nil
}

// TODO: make multiple role support
func (r *repo) GetRolesByTitle(ctx context.Context, tx pgx.Tx, role []model.Role) (uint64, error) {
	query, args, err := squirrel.Select(
		repository.RoleIdColumn,
	).
		PlaceholderFormat(squirrel.Dollar).
		From(repository.RoleTableName).
		Where(squirrel.Eq{repository.RoleTitleColumn: role}).
		ToSql()
	if err != nil {
		return 0, repository.NewInternalError(err)
	}
	var row pgx.Row
	if tx != nil {
		row = tx.QueryRow(ctx, query, args...)
	} else {
		row = r.db.QueryRow(ctx, query, args...)
	}

	var roleId uint64
	err = row.Scan(&roleId)
	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, repository.NewNotFoundError(err)
		}
		return 0, repository.NewInternalError(err)
	}
	return roleId, nil
}
