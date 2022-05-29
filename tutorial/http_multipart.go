package main

import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	reader, err := r.MultipartReader()
	if err != nil {
		log.Println("get multi part reader error: ", err.Error())
		return
	}
	metadata, err := reader.NextPart()
	if err != nil {
		log.Println("get metadata part error: ", err.Error())
		return
	}
	log.Println("metadata part name: ", metadata.FormName())
	filepart, err := reader.NextPart()
	if err != nil {
		log.Println("get file part error: ", err.Error())
		return
	}
	log.Println("file part name: ", filepart.FormName())
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
