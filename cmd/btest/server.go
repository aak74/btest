package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Store struct {
	isStopping bool
}

type Server struct {
	router *mux.Router
	store  Store
}

const (
	responseHello      = "Hello"
	responseOK         = "OK"
	responseRestarting = "Restarting"
	responseStopping   = "Stopping"
)

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
	s.router.Use(s.logRequest)
	s.router.HandleFunc("/ping", s.ping()).Methods("GET")
	s.router.HandleFunc("/hello", s.hello()).Methods("GET")
	s.router.HandleFunc("/restart", s.restart()).Methods("GET")
	s.router.HandleFunc("/stop", s.stop()).Methods("GET")
}

func (s *Server) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)
		log.Printf("%s %s %s\n", r.Method, r.RequestURI, strconv.Itoa(rw.code))
	})
}

func (s *Server) ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if s.store.isStopping {
			s.respond(w, http.StatusInternalServerError, responseStopping)
			return
		}
		s.respond(w, http.StatusOK, responseOK)
	}
}

func (s *Server) hello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, http.StatusOK, responseHello)
	}
}

func (s *Server) stop() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.store.isStopping = true
		s.respond(w, http.StatusOK, responseStopping)
	}
}

func (s *Server) restart() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.store.isStopping = false
		s.respond(w, http.StatusOK, responseRestarting)
	}
}

func (s *Server) error(w http.ResponseWriter, code int, err error) {
	s.respondJson(w, code, map[string]string{"error": err.Error()})
}

func (s *Server) respondJson(w http.ResponseWriter, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		w.Header().Add("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (s *Server) respond(w http.ResponseWriter, code int, data string) {
	w.WriteHeader(code)
	_, err := w.Write([]byte(data))
	if err != nil {
		log.Fatal(err)
	}
}
