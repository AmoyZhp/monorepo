package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"zamrepo/gomoku/executor"
)

// Handler http handler
type Handler struct {
	executor executor.Executor
}

func (imp *Handler) handler(w http.ResponseWriter, r *http.Request) {
	bodyByte, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("read body error: ", err.Error())
		return
	}
	move := executor.Move{}
	err = json.Unmarshal(bodyByte, &move)
	if err != nil {
		log.Println("json unmarshal error: ", err.Error())
		return
	}
	retMove := imp.executor.GetNextMove(move)
	retByte, _ := json.Marshal(retMove)
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	w.Write(retByte)
}

func main() {
	handler := Handler{executor: executor.NewGomokuExectuor()}
	http.HandleFunc("/get-move", handler.handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
