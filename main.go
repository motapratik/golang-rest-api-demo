package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Item - Here need to add comment for unexported
type Item struct {
	Data string `json:"mydata"`
	// If lower case data - then not accessible outside in API
	// for changine name in api we have to add separate json : name
}

var mydata []Item = []Item{}

var data []string = []string{}

func test1(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("This is Testing API"))
}

func test2(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		ID string
	}{"123"})
}
func addtoslice(w http.ResponseWriter, req *http.Request) {
	routervar := mux.Vars(req)["item"]
	data = append(data, routervar)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func addItem(w http.ResponseWriter, req *http.Request) {
	var itemdata Item
	json.NewDecoder(req.Body).Decode(&itemdata)

	mydata = append(mydata, itemdata)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mydata)
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/test1", test1)
	router.HandleFunc("/test2", test2)
	router.HandleFunc("/addtoSlice/{item}", addtoslice)
	router.HandleFunc("/addItem", addItem).Methods("POST")
	http.ListenAndServe(":8090", router)
}
