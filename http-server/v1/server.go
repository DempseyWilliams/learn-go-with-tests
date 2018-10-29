package main

import (
	"fmt"
	"net/http"
)

type PlayerStore interface {
	GetPlayerScore(name string) (int, bool)
	RecordWin(name string)
}

type PlayerServer struct {
	store PlayerStore
}

func (playServer *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	switch r.Method {
	case http.MethodPost:
		playServer.processWin(w, player)
	case http.MethodGet:
		playServer.showScore(w, player)
	}
}

func (playServer *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score, found := playServer.store.GetPlayerScore(player)
	if !found {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, score)
}

func (playServer *PlayerServer) processWin(w http.ResponseWriter, player string) {
	playServer.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}
