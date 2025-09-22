package models

type Config struct {
	App APP      `envconfig:"APP"`
	DB  Postgres `envconfig:"POSTGRES"`
}

type APP struct {
	PortAuthGRPC   string `envconfig:"PORT_GRPC_AUTH"`
	PortAccessGRPC string `envconfig:"PORT_GRPC_ACCESS"`
	PortAuthHTTP   string `envconfig:"PORT_HTTP_AUTH"`
	PortDocs       string `envconfig:"PORT_DOCS"`
}

type Postgres struct {
	User     string `envconfig:"USER" required:"true"`
	Password string `envconfig:"PASSWORD" required:"true"`
	Host     string `envconfig:"HOST" required:"true"` // localhost
	Port     string `envconfig:"PORT" required:"true"` //:5432
	Database string `envconfig:"DATABASE" required:"true"`
	SSLMode  string `envconfig:"SSL_MODE" default:"disable"`
}
