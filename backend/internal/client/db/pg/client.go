package pg

import (
	"context"

	"github.com/Krab1o/meebin/internal/client/db"
	"github.com/jackc/pgx/v5/pgxpool"
)

// there could be more connection in clients
type pgClient struct {
	masterDBC db.DB
}

func New(ctx context.Context, dsn string) (db.Client, error) {
	dbc, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	return &pgClient{
		masterDBC: &pg{dbc: dbc},
	}, nil
}

func (c *pgClient) Close() error {
	if c.masterDBC != nil {
		return c.masterDBC.Close()
	}
	return nil
}

func (c *pgClient) DB() db.DB {
	return c.masterDBC
}
