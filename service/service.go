package service

import "github.com/GabrielBG99/vxchan/domain"

type Service struct {
	boardRepository      domain.BoardRepository
	threadRepository     domain.ThreadRepository
	attachmentRepository domain.AttachmentRepository
}

func NewService(br domain.BoardRepository, tr domain.ThreadRepository, ar domain.AttachmentRepository) *Service {
	s := &Service{
		boardRepository:      br,
		threadRepository:     tr,
		attachmentRepository: ar,
	}
	return s
}
