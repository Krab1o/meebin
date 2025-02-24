package env

import (
	"fmt"
	"os"

	"github.com/Krab1o/meebin/internal/config"
)

const (
	pgHostEnvName     = "PG_HOST"
	pgPortEnvName     = "PG_PORT"
	pgUserEnvName     = "PG_USER"
	pgPasswordEnvName = "PG_PASSWORD"
	pgDatabaseEnvName = "PG_DB"
)

type pgConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func NewPGConfig() config.PGConfig {
	//TODO: Add validation
	host := os.Getenv(pgHostEnvName)
	port := os.Getenv(pgPortEnvName)
	user := os.Getenv(pgUserEnvName)
	password := os.Getenv(pgPasswordEnvName)
	dbname := os.Getenv(pgDatabaseEnvName)

	return &pgConfig{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		DBName:   dbname,
	}
}

func (c *pgConfig) DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		c.Host,
		c.Port,
		c.User,
		c.Password,
		c.DBName,
	)
}
