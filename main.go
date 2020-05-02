package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func test1(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("This is Testing API"))
}

func test2(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		ID string
	}{"123"})
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/test1", test1)
	router.HandleFunc("/test2", test2)
	http.ListenAndServe(":8090", router)
}
