package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"zamrepo/gomoku/executor"
)

// Handler http handler
type Handler struct {
	executor executor.Engine
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
	retMove, err := imp.executor.Predict(move)
	fmt.Println("predcit return error : ", err.Error())
	retByte, _ := json.Marshal(retMove)
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	w.Write(retByte)
}

func main() {
	handler := Handler{executor: executor.NewEngine()}
	http.HandleFunc("/get-move", handler.handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
