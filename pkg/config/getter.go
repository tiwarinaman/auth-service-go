package config

import "auth/config"

func Get() config.Configuration {
	return configuration
}
