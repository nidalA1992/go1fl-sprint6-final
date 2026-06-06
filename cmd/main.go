package main

import (
	"log"
	"os"

	s "github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "app-", log.Ldate|log.Ltime|log.Lshortfile|log.LstdFlags)
	server := s.NewServer(logger)

	err := server.Server.ListenAndServe()
	if err != nil {
		server.Logger.Fatal("Server failed to start: " + err.Error())
	}
}
