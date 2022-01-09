package dtos

// RecordQuery for filtering records
type RecordQuery struct {
	StartDate string `bson:"startDate" json:"startDate"`
	EndDate   string `bson:"endDate" json:"endDate"`
	MinCount  int    `bson:"minCount" json:"minCount"`
	MaxCount  int    `bson:"maxCount" json:"maxCount"`
}
