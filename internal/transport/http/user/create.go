package user

import (
	"encoding/json"
	"github.com/Akezhan1/lecvisitor/internal/manager/user"
	"net/http"
)

type createUserResponse struct {
	ID int `json:"id"`
}

func (h *handler) createUserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var userInput user.User
		if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
			h.writeResponse(w, newErrResponse(err), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		id, err := h.userSvc.CreateUser(r.Context(), userInput)
		if err != nil {
			h.writeResponse(w, newErrResponse(err), http.StatusBadRequest)
			return
		}

		h.writeResponse(w, createUserResponse{ID: id}, http.StatusCreated)
	default:
		h.writeResponse(w, newErrResponse(errMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
