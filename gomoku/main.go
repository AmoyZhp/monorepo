package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// Move chess act
type Move struct {
	Row    int `json:"row"`
	Col    int `json:"col"`
	Player int `json:"player"`
}

// Handler http handler
type Handler struct {
	engine Engine
}

func (imp *Handler) handler(w http.ResponseWriter, r *http.Request) {
	bodyByte, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("read body error: ", err.Error())
		return
	}
	move := Move{}
	err = json.Unmarshal(bodyByte, &move)
	if err != nil {
		log.Println("json unmarshal error: ", err.Error())
		return
	}
	retMove := imp.engine.Predict(move)
	retByte, _ := json.Marshal(retMove)
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	w.Write(retByte)
}

func main() {
	handler := Handler{engine: newEngine()}
	http.HandleFunc("/get-move", handler.handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
