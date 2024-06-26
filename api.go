package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteJSON(w http.ResponseWriter, data any) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(data)
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

func makeHTTPHandler(fn apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

type Server struct {
	Addr string
}

func NewServer(addr string) *Server {
	return &Server{
		Addr: addr,
	}
}

func (s *Server) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/actors", makeHTTPHandler(s.handleGetActors)).Methods("GET")

	log.Println("Server is running on", s.Addr)

	http.ListenAndServe(s.Addr, router)
}

func (s *Server) handleGetActors(w http.ResponseWriter, r *http.Request) error {
	mL := new(MetoriteLanding)

	return WriteJSON(w, mL)
}
