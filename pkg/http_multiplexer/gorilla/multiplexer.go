package gorilla

import (
	"net/http"

	gorilla_mux "github.com/gorilla/mux"
)

type multiplexer struct {
	initialized bool
	host        string
	mux         *gorilla_mux.Router
	port        int
	server      *http.Server
}
