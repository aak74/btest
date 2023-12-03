package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Store struct {
	isStopping bool
}

type Server struct {
	router *mux.Router
	store  Store
}

func NewServer() *Server {
	s := &Server{
		router: mux.NewRouter(),
		store:  Store{false},
	}

	s.configureRouter()
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) configureRouter() {
	s.router.HandleFunc("/ping", s.ping()).Methods("GET")
	s.router.HandleFunc("/hello", s.hello()).Methods("GET")
	// s.router.HandleFunc("/restart", s.restart()).Methods("GET")
	// s.router.HandleFunc("/shutdown", s.shutdown()).Methods("GET")
}

func (s *Server) ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, http.StatusOK, "OK")
	}
}

func (s *Server) hello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, http.StatusOK, "hello")
	}
}

func (s *Server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *Server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	// if data != nil {
	// 	json.NewEncoder(w).Encode(data)
	// }
}
