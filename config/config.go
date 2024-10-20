package config

type Configuration struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
}

type ServerConfig struct {
	Host string
	Port string
}

type DatabaseConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}

type JWTConfig struct {
	Secret string
	Issuer string
}
