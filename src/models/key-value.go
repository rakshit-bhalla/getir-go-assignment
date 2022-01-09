package models

// KeyValue
type KeyValue struct {
	Key   string `bson:"key" json:"key"`
	Value string `bson:"value" json:"value"`
}
