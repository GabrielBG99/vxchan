package domain

import "errors"

var (
	ErrBoardNotFound      = errors.New("Board not found")
	ErrBoardAlreadyExists = errors.New("Board already exists")
)

// Board ...
type Board struct {
	Tag         string
	Name        string
	Description string
}

type BoardRepository interface {
	// CreateBoard create a new board
	CreateBoard(tag, name, description string) (Board, error)

	// ListBoards list all boards
	ListBoards() ([]Board, error)
}
