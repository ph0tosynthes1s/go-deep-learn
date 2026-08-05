package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func GetErrorResponse(w http.ResponseWriter, handlerName string, err error, statusCode int) {
	w.WriteHeader(statusCode)
	buf := bytes.NewBufferString(handlerName)
	buf.WriteString(err.Error())
	buf.WriteString("\n")
	_, _ = w.Write(buf.Bytes())
}

type StockResponse struct {
	Stock int `json:"stock"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Header)
	body := make([]byte, r.ContentLength)
	stockResponse := &StockResponse{
		Stock: 1,
	}
	body, _ = json.Marshal(stockResponse)
	_, _ = w.Write(body)
}

// TODO
func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8089", nil))
}
