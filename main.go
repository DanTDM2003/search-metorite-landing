package main

import (
	"log"
)

func main() {
	repo, err := NewPostgresConnection()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	server := NewServer("localhost:8080", repo)
	server.Run()
}
