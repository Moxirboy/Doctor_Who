package main

import (
	"log"
	configs "testDeployment/internal/common/config"
	"testDeployment/internal/server"
)



func main() {

	var (
		config = configs.Configuration()
	)

	s := server.NewServer(config)
	log.Fatal(s.Run())
}
