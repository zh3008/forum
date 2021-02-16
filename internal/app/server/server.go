package server

import (
	"log"

	"github.com/zh3008/forum.git/internal/app/controllers"
)

//APIServer ..
type APIServer struct {
	config *Config
	router *controllers.Router
}

//New ..
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		router: controllers.NewRouter(),
	}
}

//Start ..
func (s *APIServer) Start() error {

	log.Println("Starting api server on port :8080")
	return nil
}
