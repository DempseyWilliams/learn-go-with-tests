package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Player struct {
	Name string
	Wins int
}

type PlayerStore interface {
	GetPlayerScore(name string) (int, bool)
	RecordWin(name string)
	GetLeague() []Player
}

type PlayerServer struct {
	store PlayerStore
	http.Handler
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)
	p.store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))

	p.Handler = router

	return p
}

func (playServer *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(playServer.store.GetLeague())

	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
}

func (playServer *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {

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
