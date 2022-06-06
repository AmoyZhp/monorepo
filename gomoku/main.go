package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// Move 行动
type Move struct {
	Row    int `json:"row"`
	Col    int `json:"col"`
	Player int `json:"player"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	bodyByte, err := io.ReadAll(r.Body)
	if err != nil {
		// TODO
		return
	}
	move := &Move{}
	err = json.Unmarshal(bodyByte, move)
	if err != nil {
		// TODO
		return
	}
	log.Println("move: ", move)
	retMove := &Move{
		Col:    move.Col + 1,
		Row:    move.Row + 1,
		Player: move.Player + 1,
	}
	retByte, err := json.Marshal(retMove)
	if err != nil {
		// TODO
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(retByte)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
