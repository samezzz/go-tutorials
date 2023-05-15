package storage

import "github.com/samezzz/go-postgres/types"

type Storage interface {
	Get(int) *types.Todo
}