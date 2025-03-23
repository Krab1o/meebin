package pg

import (
	"context"

	"github.com/Krab1o/meebin/internal/client/db"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	TxKey string = "tx"
)

// This could be a client already, but we wrap it again for cases when
// there is more than one connection to database
type pg struct {
	dbc *pgxpool.Pool
}

func NewDB(pool *pgxpool.Pool) db.DB {
	return &pg{
		dbc: pool,
	}
}

func (p *pg) Close() error {
	p.dbc.Close()
	return nil
}

func (p *pg) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error) {
	return p.dbc.BeginTx(ctx, txOptions)
}

func (p *pg) ExecContext(
	ctx context.Context,
	query string,
	args ...any,
) (pgconn.CommandTag, error) {
	tx, ok := ctx.Value(TxKey).(pgx.Tx)
	if ok {
		return tx.Exec(ctx, string(query), args...)
	}

	return p.dbc.Exec(ctx, string(query), args...)
}
func (p *pg) QueryContext(ctx context.Context, query string, args ...any) (pgx.Rows, error) {
	tx, ok := ctx.Value(TxKey).(pgx.Tx)
	if ok {
		return tx.Query(ctx, string(query), args...)
	}

	return p.dbc.Query(ctx, string(query), args...)
}
func (p *pg) QueryRowContext(ctx context.Context, query string, args ...any) pgx.Row {
	tx, ok := ctx.Value(TxKey).(pgx.Tx)
	if ok {
		return tx.QueryRow(ctx, string(query), args...)
	}
	return p.dbc.QueryRow(ctx, string(query), args...)
}

func (p *pg) Ping(ctx context.Context) error {
	return p.dbc.Ping(ctx)
}
