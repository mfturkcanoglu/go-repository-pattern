package main

import (
	"log"

	"github.com/mfturkcanoglu/go-repository-pattern/pkg/server"
)

var (
	logger *log.Logger = log.Default()
)

func main() {
	s := server.NewServer(logger)

	s.ListenAndServe()
}
