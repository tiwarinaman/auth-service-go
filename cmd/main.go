package main

import "auth/pkg/bootstrap"

func main() {
	serve()
}

func serve() {
	bootstrap.Serve()
}
