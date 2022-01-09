package controllers

import (
	"encoding/json"
	"net/http"

	"rakshit.dev/m/src/dtos"
	"rakshit.dev/m/src/errors"
	"rakshit.dev/m/src/services"
	"rakshit.dev/m/src/utils"
)

// RecordsController
type RecordsController interface {
	GetRecords(w http.ResponseWriter, r *http.Request)
}

type recordsController struct {
	recordsService services.RecordsService
}

// NewRecordsController
func NewRecordsController(
	recordsService services.RecordsService,
) RecordsController {
	return &recordsController{
		recordsService: recordsService,
	}
}

// Get records as per request
func (rc *recordsController) GetRecords(w http.ResponseWriter, r *http.Request) {
	var recordQuery dtos.RecordQuery
	if err := json.NewDecoder(r.Body).Decode(&recordQuery); err != nil {
		errors.InternalServerError(w)
		return
	}
	startDate, err := utils.StrToTime(recordQuery.StartDate)
	if err != nil {
		errors.BadParam(w, "startDate")
		return
	}
	endDate, err := utils.StrToTime(recordQuery.EndDate)
	if err != nil {
		errors.BadParam(w, "endDate")
		return
	}
	records, err := rc.recordsService.GetRecords(
		startDate,
		endDate,
		recordQuery.MinCount,
		recordQuery.MaxCount,
	)
	if err != nil {
		errors.InternalServerError(w)
		return
	}
	recordsRes := dtos.RecordResponse{
		Code:    0,
		Msg:     "Success",
		Records: records,
	}
	jsonBytes, err := json.Marshal(recordsRes)
	if err != nil {
		errors.InternalServerError(w)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
