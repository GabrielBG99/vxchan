package fs

import "os"

type service struct {
	folder string
}

func NewService(folder string) (*service, error) {
	if err := os.Mkdir(folder, 0755); err != nil && !os.IsExist(err) {
		return nil, err
	}

	s := &service{
		folder: folder,
	}
	return s, nil
}
