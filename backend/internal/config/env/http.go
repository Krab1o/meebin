package env

import (
	"net"
	"os"

	"github.com/Krab1o/meebin/internal/config"
)

const (
	httpHostEnvName = "HTTP_HOST"
	httpPortEnvName = "HTTP_PORT"
)

type httpConfig struct {
	host string
	port string
}

func NewHTTPConfig() (config.HTTPConfig, error) {
	host := os.Getenv(httpHostEnvName)
	port := os.Getenv(httpPortEnvName)

	return &httpConfig{
		host: host,
		port: port,
	}, nil
}

func (c *httpConfig) Port() string {
	return c.port
}

func (c *httpConfig) Address() string {
	return net.JoinHostPort(c.host, c.port)
}
