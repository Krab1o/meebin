package env

import (
	"errors"
	"os"
	"strconv"

	"github.com/Krab1o/meebin/internal/config"
)

const (
	jwtSecretEnvName         = "JWT_SECRET"
	jwtAccessTimeoutEnvName  = "JWT_ACCESS_TIMEOUT"
	jwtRefreshTimeoutEnvName = "JWT_REFRESH_TIMEOUT"
	jwtTimeoutParseError     = "Unable to parse JWT timeout"
)

type jwtConfig struct {
	jwtSecret      []byte
	accessTimeout  int
	refreshTimeout int
}

// TODO: add error messages
// TODO: move error messages to one place for config
func NewJWTConfig() (config.JWTConfig, error) {
	jwt := os.Getenv(jwtSecretEnvName)
	if len(jwt) == 0 {
		return nil, errors.New("No env var")
	}
	accessTime := os.Getenv(jwtAccessTimeoutEnvName)
	if len(accessTime) == 0 {
		return nil, errors.New("No env var")
	}
	refreshTime := os.Getenv(jwtRefreshTimeoutEnvName)
	if len(refreshTime) == 0 {
		return nil, errors.New("No env var")
	}
	accessTimeVal, err := strconv.Atoi(accessTime)
	if err != nil {
		return nil, errors.New(jwtTimeoutParseError)
	}
	refreshTimeVal, err := strconv.Atoi(refreshTime)
	if err != nil {
		return nil, errors.New(jwtTimeoutParseError)
	}
	return &jwtConfig{
		jwtSecret:      []byte(jwt),
		accessTimeout:  accessTimeVal,
		refreshTimeout: refreshTimeVal,
	}, nil
}

func (c *jwtConfig) Secret() []byte {
	return c.jwtSecret
}

func (c *jwtConfig) AccessTimeout() int {
	return c.accessTimeout
}

func (c *jwtConfig) RefreshTimeout() int {
	return c.accessTimeout
}
