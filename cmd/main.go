package main

import (
	"github.com/cocvu99/sabaody-kv/internal/config"
	"github.com/cocvu99/sabaody-kv/internal/server"
)

func main() {

	server.Start(config.Port)
}
