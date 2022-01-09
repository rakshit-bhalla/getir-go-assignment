package dtos

import "rakshit.dev/m/src/models"

type RecordResponse struct {
	Code    int             `bson:"code" json:"code"`
	Msg     string          `bson:"msg" json:"msg"`
	Records []models.Record `bson:"records" json:"records"`
}
