package models

import "time"

// Record
type Record struct {
	Key        string    `bson:"key" json:"key"`
	TotalCount int       `bson:"totalCount" json:"totalCount"`
	CreatedAt  time.Time `bson:"createdAt" json:"createdAt"`
}
