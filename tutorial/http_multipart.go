package main

import (
	"encoding/json"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

// MetaData http meta-data form 携带的元数据
type MetaData struct {
	RequestID string `json:"request_id"`
}

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
	readMetaData(metadata)
	log.Println("metadata part name: ", metadata.FormName())
	filepart, err := reader.NextPart()
	if err != nil {
		log.Println("get file part error: ", err.Error())
		return
	}
	log.Println("file part name: ", filepart.FormName())
	readFile(filepart)
	w.Write([]byte("request successfule"))
}

func readMetaData(part *multipart.Part) {
	dataCache := make([]byte, 1024)
	n, err := part.Read(dataCache)
	if err != nil && err != io.EOF {
		log.Println("read meta data error. ", err.Error())
		return
	}
	data := &MetaData{}
	err = json.Unmarshal(dataCache[:n], data)
	if err != nil {
		log.Println("json unmarshal error. ", err.Error())
		return
	}
	log.Printf("read meta data %+v", data)
}

func readFile(part *multipart.Part) {
	f, err := os.Create("temp.png")
	if err != nil {
		log.Println("create file error. ", err.Error())
		return
	}
	buf := make([]byte, 1024)
	for {
		n, err := part.Read(buf)
		if err != nil && err != io.EOF {
			log.Println("read file data error. ", err.Error())
			return
		}
		if err == io.EOF || n == 0 {
			break
		}
		_, err = f.Write(buf[:n])
		if err != nil {
			log.Println("write file error. ", err.Error())
			return
		}
	}
	log.Println("write file success. ")
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
