package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", MessageHandler)

	log.Println(http.ListenAndServe("127.0.0.1:8080", r))
}

func MessageHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello from server 1"))
	if err != nil {
		log.Println(err)
	}

}
