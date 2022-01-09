package controllers

import (
	"encoding/json"
	"net/http"

	"rakshit.dev/m/src/errors"
	"rakshit.dev/m/src/models"
	"rakshit.dev/m/src/services"
)

// InMemoryController
type InMemoryController interface {
	Write(w http.ResponseWriter, r *http.Request)
	Read(w http.ResponseWriter, r *http.Request)
}

type inMemoryController struct {
	inMemoryService services.InMemoryService
}

// NewInMemoryController
func NewInMemoryController(
	inMemoryService services.InMemoryService,
) InMemoryController {
	return &inMemoryController{
		inMemoryService: inMemoryService,
	}
}

// Takes  key-value pair from body and saves it into the db and sends it back as a response
func (imc *inMemoryController) Write(w http.ResponseWriter, r *http.Request) {
	var keyValue models.KeyValue
	if err := json.NewDecoder(r.Body).Decode(&keyValue); err != nil {
		errors.InternalServerError(w)
		return
	}
	savedKeyValue, err := imc.inMemoryService.Write(keyValue)
	if err != nil {
		errors.InternalServerError(w)
		return
	}
	jsonBytes, err := json.Marshal(savedKeyValue)
	if err != nil {
		errors.InternalServerError(w)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

// Takes key from query param and sends key-value pair as a response
func (imc *inMemoryController) Read(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	savedKeyValue, err := imc.inMemoryService.Read(key)
	if err != nil {
		errors.InternalServerError(w)
		return
	}
	if savedKeyValue == nil {
		errors.NotFound(w)
		return
	}
	jsonBytes, err := json.Marshal(savedKeyValue)
	if err != nil {
		errors.InternalServerError(w)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
