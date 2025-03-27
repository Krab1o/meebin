package env

import (
	"errors"
	"fmt"
	"os"

	"github.com/Krab1o/meebin/internal/config"
)

const (
	pgPortEnvName     = "PG_PORT"
	pgUserEnvName     = "PG_USER"
	pgPasswordEnvName = "PG_PASSWORD"
	pgDatabaseEnvName = "PG_DB"
	pgHostEnvName     = "DB_HOST"
)

type pgConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func NewPGConfig() (config.PGConfig, error) {

	host := os.Getenv(pgHostEnvName)
	if len(host) == 0 {
		return nil, errors.New(fmt.Sprintf("empty %s", pgHostEnvName))
	}
	port := os.Getenv(pgPortEnvName)
	if len(host) == 0 {
		return nil, errors.New(fmt.Sprintf("empty %s", pgPortEnvName))
	}
	user := os.Getenv(pgUserEnvName)
	if len(host) == 0 {
		return nil, errors.New(fmt.Sprintf("empty %s", pgUserEnvName))
	}
	password := os.Getenv(pgPasswordEnvName)
	if len(host) == 0 {
		return nil, errors.New(fmt.Sprintf("empty %s", pgPasswordEnvName))
	}
	dbname := os.Getenv(pgDatabaseEnvName)
	if len(host) == 0 {
		return nil, errors.New(fmt.Sprintf("empty %s", pgDatabaseEnvName))
	}

	return &pgConfig{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		DBName:   dbname,
	}, nil
}

func (c *pgConfig) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		c.Host,
		c.Port,
		c.User,
		c.Password,
		c.DBName,
	)
}
