package main

import (
	"log"
	"net/http"
)

type InMemroyPlayerScore struct {
}

func (i *InMemroyPlayerScore) GetPlayerScore(name string) int {
	return 123
}
func (i *InMemroyPlayerScore) RecordWin(name string) {}

func main() {

	server := &PlayerServer{&InMemroyPlayerScore{}}
	if err := http.ListenAndServe(":8000", server); err != nil {
		log.Fatalf("could not listen on port :8000 %v", err)
	}
}
