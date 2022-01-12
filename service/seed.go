package service

import (
	"errors"

	"github.com/GabrielBG99/vxchan/domain"
)

func (s Service) Seed() error {
	if _, err := s.CreateBoard(
		"b",
		"Random",
		"The stories and information posted here are artistic works of fiction and falsehood. Only a fool would take anything posted here as fact.",
	); err != nil && !errors.Is(err, domain.ErrBoardAlreadyExists) {
		return err
	}

	return nil
}
