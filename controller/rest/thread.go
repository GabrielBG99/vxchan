package rest

import (
	"mime/multipart"
	"net/http"
	"strconv"
	"time"

	"github.com/GabrielBG99/vxchan/service"
	"github.com/go-chi/chi"
)

type (
	ReplyRequest struct {
		File    multipart.File `multipart:"file"`
		Comment string         `multipart:"comment"`
	}

	ThreadRequest struct {
		File        multipart.File `multipart:"file"`
		Description string         `multipart:"description"`
	}

	ReplyResponse struct {
		ID         int64     `json:"id"`
		CreatedAt  time.Time `json:"createdAt"`
		Content    string    `json:"content"`
		Attachment string    `json:"attachment"`
	}

	ThreadResponse struct {
		ID          int64           `json:"id"`
		CreatedAt   time.Time       `json:"createdAt"`
		BoardID     string          `json:"boardID"`
		Description string          `json:"description"`
		Attachment  string          `json:"attachment"`
		Replies     []ReplyResponse `json:"replies"`
	}
)

func replyFromService(r service.Reply) ReplyResponse {
	return ReplyResponse{
		ID:         r.ID,
		CreatedAt:  r.CreatedAt,
		Content:    r.Content,
		Attachment: r.Attachment,
	}
}

func threadFromService(t service.Thread) ThreadResponse {
	replies := make([]ReplyResponse, 0)
	for _, r := range t.Replies {
		replies = append(replies, replyFromService(r))
	}

	return ThreadResponse{
		ID:          t.ID,
		CreatedAt:   t.CreatedAt,
		BoardID:     t.BoardID,
		Description: t.Description,
		Attachment:  t.Attachment,
		Replies:     replies,
	}
}

func v1CreateThread(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		boardID := chi.URLParam(r, "boardID")

		var body ThreadRequest
		if err := DecodeMultipartForm(r, &body); err != nil {
			JSONResponse(w, http.StatusInternalServerError, ErrInternal)
			return
		}
		defer body.File.Close()

		t, err := svc.CreateThread(boardID, body.Description, body.File)
		if err != nil {
			JSONResponse(w, http.StatusInternalServerError, ErrInternal)
			return
		}

		JSONResponse(w, http.StatusCreated, threadFromService(t))
	}
}

func v1ListThreads(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		boardID := chi.URLParam(r, "boardID")

		query := r.URL.Query()

		page, err := StringToInt(query.Get("page"), 0)
		if err != nil {
			JSONResponse(w, http.StatusUnprocessableEntity, ErrInvalidQueryParam)
			return
		}

		nItems, err := StringToInt(query.Get("nItems"), 10)
		if err != nil {
			JSONResponse(w, http.StatusUnprocessableEntity, ErrInvalidQueryParam)
			return
		}

		ts, err := svc.ListThreads(boardID, page, nItems)
		if err != nil {
			JSONResponse(w, http.StatusInternalServerError, ErrInternal)
			return
		}

		threads := make([]ThreadResponse, 0)
		for _, t := range ts {
			threads = append(threads, threadFromService(t))
		}

		JSONResponse(w, http.StatusOK, threads)
	}
}

func v1ReplyThread(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		threadID, err := strconv.ParseInt(chi.URLParam(r, "threadID"), 10, 64)
		if err != nil || threadID < 0 {
			JSONResponse(w, http.StatusBadRequest, ErrInvalidPathParam)
			return
		}

		var body ReplyRequest
		if err := DecodeMultipartForm(r, &body); err != nil {
			JSONResponse(w, http.StatusInternalServerError, ErrInternal)
			return
		}
		defer body.File.Close()

		rp, err := svc.ReplyThread(threadID, body.Comment, body.File)
		if err != nil {
			JSONResponse(w, http.StatusInternalServerError, ErrInternal)
			return
		}

		JSONResponse(w, http.StatusCreated, replyFromService(rp))
	}
}
