package postgresql

import (
	"github.com/GabrielBG99/vxchan/domain"
	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

type Board struct {
	gorm.Model
	Tag         string `gorm:"uniqueIndex"`
	Name        string
	Description string
}

func (b Board) ToDomain() domain.Board {
	return domain.Board{
		Tag:         b.Tag,
		Name:        b.Name,
		Description: b.Description,
	}
}

func (c connector) initBoard() error {
	return c.db.AutoMigrate(&Board{})
}

func (c connector) CreateBoard(tag, name, description string) (domain.Board, error) {
	b := Board{
		Tag:         tag,
		Name:        name,
		Description: description,
	}

	r := c.db.Create(&b)
	if err := r.Error; err != nil {
		if e, ok := err.(*pgconn.PgError); ok && alreadyExistsErrorPattern.MatchString(e.Detail) {
			return domain.Board{}, domain.ErrBoardAlreadyExists
		}

		return domain.Board{}, err
	}

	board := b.ToDomain()
	return board, nil
}

func (c connector) ListBoards() ([]domain.Board, error) {
	var data []Board
	r := c.db.Model(&Board{}).Find(&data)
	if err := r.Error; err != nil {
		return nil, err
	}

	boards := make([]domain.Board, 0)
	for _, b := range data {
		boards = append(boards, b.ToDomain())
	}

	return boards, nil
}
