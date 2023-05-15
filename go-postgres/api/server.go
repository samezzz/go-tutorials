package api

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/samezzz/go-postgres/storage"
)

type Server struct {
	listenAddress string
	store         storage.Storage
}

func NewServer(listenAddress string, store storage.Storage) *Server {
	return &Server{
		listenAddress: listenAddress,
		store:         store,
	}
}

func (s *Server) Start() error {
	router := mux.NewRouter()

	router.HandleFunc("/todos", s.handleGetAllTodos).Methods("GET")
	router.HandleFunc("/todo/{id}", s.handleGetTodo).Methods("GET")
	router.HandleFunc("/todo", s.handleCreateTodo).Methods("POST")
	router.HandleFunc("/todo/{id}", s.handleUpdateTodo).Methods("PUT")
	router.HandleFunc("/todo/{id}", s.handleDeleteTodo).Methods("DELETE")

	return http.ListenAndServe(s.listenAddress, nil)
}
