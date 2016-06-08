package api

import (
	"net/http"

	"github.com/bcrusu/lunch/backEnd-go/api/handlers"
	"github.com/bcrusu/lunch/backEnd-go/api/middleware"
	"goji.io"
)

func NewMux() http.Handler {
	mux := goji.NewMux()

	mux.Use(middleware.NewCorsHandler())
	mux.UseC(middleware.ErrorHandler)

	//TODO: AUTH
	//TODO: auto save changes middleware

	handlers.ConfigureMux(mux)

	return mux
}
