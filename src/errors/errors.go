package errors

import (
	"net/http"
)

func InternalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("internal server error"))
}

func NotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("not found"))
}

func BadParam(w http.ResponseWriter, p string) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("invalid param: " + p))
}
