package http

import (
	"github.com/Akezhan1/lecvisitor/internal/manager/user"
	userhttp "github.com/Akezhan1/lecvisitor/internal/transport/http/user"
	"net/http"
)

func NewHandlersMux(userSvc user.Service) *http.ServeMux {
	userHandle := userhttp.NewHandler(userSvc)

	mux := http.NewServeMux()
	mux.Handle("/user", userHandle.ServeMux())

	return mux
}
