package rest

import (
	"net/http"

	"github.com/GabrielBG99/vxchan/service"
)

type BoardResponse struct {
	Tag         string `json:"tag"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func boardFromService(b service.Board) BoardResponse {
	return BoardResponse{
		Tag:         b.Tag,
		Name:        b.Name,
		Description: b.Description,
	}
}

func v1ListBoards(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bs, err := svc.ListBoards()
		if err != nil {
			JSONResponse(w, http.StatusInternalServerError, ErrInternal)
			return
		}

		boards := make([]BoardResponse, 0)
		for _, b := range bs {
			boards = append(boards, boardFromService(b))
		}

		JSONResponse(w, http.StatusOK, boards)
	}
}
