package fs

import (
	"fmt"
	"io"
	"os"
	"path"

	"github.com/GabrielBG99/vxchan/domain"
	"github.com/google/uuid"
)

func (s service) SaveAttachment(fileData io.Reader) (domain.Attachment, error) {
	filename := fmt.Sprintf("%s.jpg", uuid.New().String())

	f, err := os.Create(path.Join(s.folder, filename))
	if err != nil {
		return "", err
	}
	defer f.Close()

	_, err = io.Copy(f, fileData)
	if err != nil {
		return "", err
	}

	return domain.Attachment(filename), nil
}
