package rest

import (
	"encoding/json"
	"net/http"

	"github.com/GabrielBG99/vxchan/logger"
)

func JSONResponse(w http.ResponseWriter, status int, body interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if status != http.StatusOK {
		w.WriteHeader(status)
	}

	if err := json.NewEncoder(w).Encode(body); err != nil {
		logger.GetLogger().Error(err)
		return
	}
}
