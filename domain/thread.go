package domain

import (
	"errors"
	"time"
)

var (
	ErrThreadNotFound = errors.New("Thread not found")
)

type (
	// Reply ...
	Reply struct {
		ID         int64
		CreatedAt  time.Time
		ThreadID   int64
		Content    string
		Attachment Attachment
	}

	// Thread ...
	Thread struct {
		ID          int64
		CreatedAt   time.Time
		BoardID     string
		Description string
		Attachment  Attachment
		Replies     []Reply
	}
)

type ThreadRepository interface {
	// CreateThread create a new thread for a given board
	CreateThread(boardID, content string, attachment Attachment) (Thread, error)

	// ListThreads lists board's threads
	ListThreads(boardID string, page, nItems int) ([]Thread, error)

	// ReplyThread create a reply for a thread
	ReplyThread(threadID int64, comment string, attachment Attachment) (Reply, error)
}
