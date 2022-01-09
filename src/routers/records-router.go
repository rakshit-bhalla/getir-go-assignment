package routers

import (
	"net/http"

	"rakshit.dev/m/src/controllers"
	"rakshit.dev/m/src/errors"
	"rakshit.dev/m/src/utils"
)

// RecordsRouter
type RecordsRouter interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type recordsRouter struct {
	recordsController controllers.RecordsController
}

// NewRecordsRouter
func NewRecordsRouter(
	recordsController controllers.RecordsController,
) RecordsRouter {
	return &recordsRouter{
		recordsController: recordsController,
	}
}

// ServeHTTP calls controller method with respect to HTTP method and HTTP URL
func (rr *recordsRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
	switch {
	case r.Method == http.MethodPost && r.URL.Path == utils.RECORDS_BASE_URL:
		rr.recordsController.GetRecords(w, r)
		return
	default:
		errors.NotFound(w)
		return
	}
}
