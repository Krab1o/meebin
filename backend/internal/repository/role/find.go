package role

import (
	"context"
	"fmt"

	"github.com/Krab1o/meebin/internal/model"
	rep "github.com/Krab1o/meebin/internal/repository"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func (r *repo) GetUserRolesById(
	ctx context.Context,
	userId uint64,
) ([]model.Role, error) {
	query, args, err := sq.Select(
		rep.RoleTitleColumn,
	).
		PlaceholderFormat(sq.Dollar).
		From(rep.UserRoleTableName).
		LeftJoin(fmt.Sprintf("%s ON %s = %s",
			rep.RoleTableName,
			rep.Col(rep.UserRoleTableName, rep.UserRoleIdRoleColumn),
			rep.Col(rep.RoleTableName, rep.RoleIdColumn),
		),
		).Where(sq.Eq{rep.Col(rep.UserRoleTableName, rep.UserRoleIdUserColumn): userId}).
		ToSql()
	if err != nil {
		return nil, rep.NewInternalError(err)
	}

	rows, err := r.db.DB().QueryContext(ctx, query, args...)
	if err != nil {
		return nil, rep.NewInternalError(err)
	}

	var roles []model.Role
	var role model.Role
	for rows.Next() {
		err = rows.Scan(&role)
		if err != nil {
			return nil, rep.NewInternalError(err)
		}
		roles = append(roles, role)
	}

	return roles, nil
}

// TODO: make multiple role support
func (r *repo) GetRolesByTitle(ctx context.Context, role []model.Role) (uint64, error) {
	query, args, err := sq.Select(
		rep.RoleIdColumn,
	).
		PlaceholderFormat(sq.Dollar).
		From(rep.RoleTableName).
		Where(sq.Eq{rep.RoleTitleColumn: role}).
		ToSql()
	if err != nil {
		return 0, rep.NewInternalError(err)
	}
	row := r.db.DB().QueryRowContext(ctx, query, args...)

	var roleId uint64
	err = row.Scan(&roleId)
	if err != nil {
		if err == pgx.ErrNoRows {
			return 0, rep.NewNotFoundError(err)
		}
		return 0, rep.NewInternalError(err)
	}
	return roleId, nil
}
