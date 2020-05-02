package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func test(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("This is Testing API"))
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/test", test)

	http.ListenAndServe(":8090", router)
}
