package domain

import (
	"errors"
	"io"
)

var (
	ErrInvalidFileType = errors.New("Invalid file type")
)

// Attachment ...
type Attachment string

type AttachmentRepository interface {
	// SaveAttachment save a file and returns a download link
	SaveAttachment(fileData io.Reader) (Attachment, error)
}
