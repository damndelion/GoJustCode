package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", MessageHandler2)
	log.Println(http.ListenAndServe("127.0.0.1:9090", r))
}

func MessageHandler2(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello from server 2"))

	if err != nil {
		log.Print(err)
	}
}
