package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
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
	repo Repository
}

func NewServer(
	addr string,
	// repo Repository,
) *Server {
	return &Server{
		Addr: addr,
		// repo: repo,
	}
}

func (s *Server) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/meteor-landings", makeHTTPHandler(s.handleGetMeteorLandings)).Methods("GET")

	log.Println("Server is running on", s.Addr)

	http.ListenAndServe(s.Addr, router)
}

func (s *Server) handleGetMeteorLandings(w http.ResponseWriter, r *http.Request) error {
	// mL := new(MetoriteLanding)
	mLs, err := s.repo.GetMetoriteLandings()
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, mLs)
}
