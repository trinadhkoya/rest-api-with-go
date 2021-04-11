package http

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	muxDispatcher = mux.NewRouter()
)

type muxRouter struct {
}

func (m muxRouter) PUT(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("PUT")

}

func (m muxRouter) DELETE(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("DELETE")

}

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (m muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}

func (m muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("POST")
}

func (m muxRouter) SERVE(port string) {
	fmt.Printf("Mux HTTP server running on port %v", port)
	http.ListenAndServe(port, muxDispatcher)

}
