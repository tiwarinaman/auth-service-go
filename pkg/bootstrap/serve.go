package bootstrap

import (
	"auth/pkg/config"
	"auth/pkg/database"
	"auth/pkg/routing"
)

func Serve() {

	// set configuration
	config.Set()

	// connect database
	database.Connect()

	// initialize router
	routing.Init()

	// register routes
	routing.RegisterRoutes()

	// serve
	routing.Serve()

}
