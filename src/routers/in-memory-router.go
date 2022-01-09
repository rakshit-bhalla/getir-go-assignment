package routers

import (
	"net/http"

	"rakshit.dev/m/src/controllers"
	"rakshit.dev/m/src/errors"
	"rakshit.dev/m/src/utils"
)

// InMemoryRouter
type InMemoryRouter interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type inMemoryRouter struct {
	inMemoryController controllers.InMemoryController
}

// NewInMemoryRouter
func NewInMemoryRouter(
	inMemoryController controllers.InMemoryController,
) InMemoryRouter {
	return &inMemoryRouter{
		inMemoryController: inMemoryController,
	}
}

// ServeHTTP calls controller method with respect to HTTP method and HTTP URL
func (imr *inMemoryRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodPost && r.URL.Path == utils.IN_MEMORY_BASE_URL:
		imr.inMemoryController.Write(w, r)
		return
	case r.Method == http.MethodGet && r.URL.Path == utils.IN_MEMORY_BASE_URL:
		imr.inMemoryController.Read(w, r)
		return
	default:
		errors.NotFound(w)
		return
	}
}
