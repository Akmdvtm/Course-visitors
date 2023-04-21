package user

import (
	"encoding/json"
	"github.com/Akezhan1/lecvisitor/internal/manager/user"
	"net/http"
)

type handler struct {
	userSvc user.Service
}

func NewHandler(userSvc user.Service) *handler {
	return &handler{
		userSvc: userSvc,
	}
}

func (h *handler) ServeMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/create", h.createUserHandler)

	return mux
}

func (h *handler) writeResponse(w http.ResponseWriter, data any, code int) {
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
