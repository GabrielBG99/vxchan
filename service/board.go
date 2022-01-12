package service

import "github.com/GabrielBG99/vxchan/domain"

type Board struct {
	Tag         string
	Name        string
	Description string
}

func boardFromDomain(b domain.Board) Board {
	return Board{
		Tag:         b.Tag,
		Name:        b.Name,
		Description: b.Description,
	}
}

func (s Service) CreateBoard(tag, name, description string) (Board, error) {
	b, err := s.boardRepository.CreateBoard(tag, name, description)
	if err != nil {
		return Board{}, err
	}

	board := boardFromDomain(b)

	return board, nil
}

func (s Service) ListBoards() ([]Board, error) {
	bs, err := s.boardRepository.ListBoards()
	if err != nil {
		return nil, err
	}

	boards := make([]Board, 0)
	for _, b := range bs {
		boards = append(boards, boardFromDomain(b))
	}

	return boards, nil
}
