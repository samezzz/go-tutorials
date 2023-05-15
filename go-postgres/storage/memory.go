package storage

import "github.com/samezzz/go-postgres/types"

type MemoryStorage struct{}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{}
}

func (s *MemoryStorage) Get(id int) *types.Todo {
	return &types.Todo{
		ID:         1,
		Title:      "Todo 1",
		IsFinished: false,
	}
}
