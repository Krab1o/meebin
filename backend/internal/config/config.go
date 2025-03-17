package config

type PGConfig interface {
	DSN() string
}

type HTTPConfig interface {
	// Address() string
	Port() string
}

type JWTConfig interface {
	Secret() []byte
	AccessTimeout() int
	RefreshTimeout() int
}
