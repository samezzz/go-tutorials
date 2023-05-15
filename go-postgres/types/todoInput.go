package types

type TodoInput struct {
	Title      string `json:"title" validate:"required"`
	IsFinished bool   `json:"isFinished" validate:"required"`
}
