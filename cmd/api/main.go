package main

import (
	"log"

	"github.com/zh3008/forum.git/internal/app/server"
)

func main() {
	config := server.NewConfig()
	s := server.New(config)
	if err := s.Start(); err != nil {
		log.Fatal((err))

	}
}
