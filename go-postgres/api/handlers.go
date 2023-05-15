package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"

	"github.com/samezzz/go-postgres/storage"
	"github.com/samezzz/go-postgres/types"
	"github.com/samezzz/go-postgres/utils"
)

func (s *Server) handleGetAllTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var todos []types.Todo
	storage.DB.Find(&todos)

	json.NewEncoder(w).Encode(todos)
}

func (s *Server) handleGetTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var todo types.Todo

	if err := storage.DB.Where("id = ?", id).First(&todo).Error; err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Todo not found")
		return
	}

	json.NewEncoder(w).Encode(todo)
}

func (s *Server) handleCreateTodo(w http.ResponseWriter, r *http.Request) {
	var input types.TodoInput

	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &input)

	validate = validator.New()
	err := validate.Struct(input)

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Validation Error")
		return
	}

	todo := &types.Todo{
		Title:      input.Title,
		IsFinished: input.IsFinished,
	}

	storage.DB.Create(todo)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(todo)
}

func (s *Server) handleUpdateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var todo types.Todo

	if err := storage.DB.Where("id = ?", id).First(&todo).Error; err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Todo not found")
		return
	}

	var input TodoInput

	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &input)

	validate = validator.New()
	err := validate.Struct(input)

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Validation Error")
		return
	}

	todo.Title = input.Title
	todo.IsFinished = input.IsFinished

	storage.DB.Save(&todo)

	json.NewEncoder(w).Encode(todo)
}

func (s *Server) handleDeleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var todo types.Todo

	if err := storage.DB.Where("id = ?", id).First(&todo).Error; err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Todo not found")
		return
	}

	storage.DB.Delete(&todo)

	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(todo)
}
