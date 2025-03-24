package db

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type Client interface {
	DB() DB
	Close() error
}

type Handler func(ctx context.Context) error

type DB interface {
	SQLExecer
	Pinger
	Transactor
	Close() error
}

type SQLExecer interface {
	// ScanExecer
	QueryExecer
}

// type ScanExecer interface {
// }

type QueryExecer interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) pgx.Row
}

type Pinger interface {
	Ping(ctx context.Context) error
}
