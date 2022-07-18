package handlers

import (
	"github.com/koopa0/go-api/internal/helpers"
	"net/http"
)

func Broker(w http.ResponseWriter, r *http.Request) {

	payload := helpers.JsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	_ = helpers.WriteJSON(w, http.StatusOK, payload)
}
