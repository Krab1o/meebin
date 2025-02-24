package main

import (
	"github.com/Krab1o/meebin/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

type serviceProvider struct {
	pgPool *pgxpool.Pool
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	return nil
}
