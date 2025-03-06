package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func postgresInit(DSN string) *pgxpool.Pool {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, DSN)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer pool.Close()

	err = pool.Ping(ctx)
	if err != nil {
		log.Fatalf("DB is not reachable, %v", err)
	}
	return pool
}
