package api

import (
	"encoding/json"
	"net/http"

	"github.com/samezzz/go-postgres/utils"
)

func (s *Server) handleGetUserByID(w http.ResponseWriter, r *http.Request) {
	user := s.store.Get(10)

	_ = utils.Round2Dec(10.3444)

	json.NewEncoder(w).Encode(user)
}

func (s *Server) handleFoo(w http.ResponseWriter, r *http.Request)    {}
func (s *Server) handleBar(w http.ResponseWriter, r *http.Request)    {}
func (s *Server) handleBaz(w http.ResponseWriter, r *http.Request)    {}
func (s *Server) handleDelete(w http.ResponseWriter, r *http.Request) {}
