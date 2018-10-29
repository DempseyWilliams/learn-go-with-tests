package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// func NewInMemoryPlayerStore() *InMemoryPlayerStore {
// 	return &InMemoryPlayerStore{map[string]int{}}
// }

// type InMemoryPlayerStore struct {
// 	store map[string]int
// }

// func (i *InMemoryPlayerStore) GetPlayerScore(name string) (int, bool) {
// 	score, found := i.store[name]
// 	return score, found
// }

// func (i *InMemoryPlayerStore) RecordWin(name string) {
// 	i.store[name]++
// }

type JSONMemoryStore struct {
	Scores map[string]int
}

// type JSONMemoryStore struct {
// 	scores TempStruct
// }

// type TempStruct struct {
// 	Pepper     int
// 	Floyd      int
// 	TestPlayer int
// }

const filename = "./jsonData.json"

func (store *JSONMemoryStore) GetPlayerScore(name string) (int, bool) {
	//read from JSON, parse it & return score/found
	dataMap := getJSONScores()

	score, found := dataMap.Scores[name]
	return score, found
}

func (store *JSONMemoryStore) RecordWin(name string) {
	// write to the JSON
	dataMap := getJSONScores()

	if _, ok := dataMap.Scores[name]; ok {
		dataMap.Scores[name]++
	} else {
		dataMap.Scores[name] = 1
	}

	saveJSONScores(dataMap)
}

func getJSONScores() *JSONMemoryStore {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("could not read data file", err)
	}

	var dataMap *JSONMemoryStore
	err = json.Unmarshal(data, &dataMap)
	if err != nil {
		log.Fatal("failure unmarshalling file", err)
	}

	return dataMap
}

func saveJSONScores(scores *JSONMemoryStore) {
	data, err := json.Marshal(scores)
	if err != nil {
		log.Fatal("failure marshalling json", err)
	}

	err = ioutil.WriteFile(filename, data, 0777)
	if err != nil {
		log.Fatal("could not write the json file", err)
	}

}

func main() {

	server := &PlayerServer{&JSONMemoryStore{}}
	// server := &PlayerServer{NewInMemoryPlayerStore()}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
