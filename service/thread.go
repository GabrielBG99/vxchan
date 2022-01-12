package service

import (
	"io"
	"time"

	"github.com/GabrielBG99/vxchan/domain"
)

type (
	Reply struct {
		ID         int64
		CreatedAt  time.Time
		ThreadID   int64
		Content    string
		Attachment string
	}

	Thread struct {
		ID          int64
		CreatedAt   time.Time
		BoardID     string
		Description string
		Attachment  string
		Replies     []Reply
	}
)

func replyFromDomain(r domain.Reply) Reply {
	return Reply{
		ID:         r.ID,
		CreatedAt:  r.CreatedAt,
		ThreadID:   r.ThreadID,
		Content:    r.Content,
		Attachment: string(r.Attachment),
	}
}

func threadFromDomain(t domain.Thread) Thread {
	replies := make([]Reply, 0)
	for _, r := range t.Replies {
		replies = append(replies, replyFromDomain(r))
	}

	return Thread{
		ID:          t.ID,
		CreatedAt:   t.CreatedAt,
		BoardID:     t.BoardID,
		Description: t.Description,
		Attachment:  string(t.Attachment),
		Replies:     replies,
	}
}

func (s Service) CreateThread(boardID, content string, attachment io.Reader) (Thread, error) {
	f, err := s.attachmentRepository.SaveAttachment(attachment)
	if err != nil {
		return Thread{}, err
	}

	t, err := s.threadRepository.CreateThread(boardID, content, f)
	if err != nil {
		return Thread{}, err
	}

	thread := threadFromDomain(t)

	return thread, nil
}

func (s Service) ListThreads(boardID string, page, nItems int) ([]Thread, error) {
	ts, err := s.threadRepository.ListThreads(boardID, page, nItems)
	if err != nil {
		return nil, err
	}

	threads := make([]Thread, 0)
	for _, t := range ts {
		threads = append(threads, threadFromDomain(t))
	}

	return nil, nil
}

func (s Service) ReplyThread(threadID int64, comment string, attachment io.Reader) (Reply, error) {
	f, err := s.attachmentRepository.SaveAttachment(attachment)
	if err != nil {
		return Reply{}, err
	}

	r, err := s.threadRepository.ReplyThread(threadID, comment, f)
	if err != nil {
		return Reply{}, err
	}

	reply := replyFromDomain(r)

	return reply, nil
}
