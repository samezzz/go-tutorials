package api

import (
	"net/http"
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
	http.HandleFunc("/user", s.handleGetUserByID)
	http.HandleFunc("/foo", s.handleFoo)
	http.HandleFunc("/bar", s.handleBar)
	http.HandleFunc("/baz", s.handleBaz)
	return http.ListenAndServe(s.listenAddress, nil)
}
