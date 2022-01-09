package repositories

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/mongo"
	"rakshit.dev/m/src/db"
	"rakshit.dev/m/src/models"
)

// RecordsRepository
type RecordsRepository interface {
	GetRecords(
		startDate time.Time, endDate time.Time, minCount int, maxCount int,
	) ([]models.Record, error)
}

type recordsRepository struct {
	recordsCollection *mongo.Collection
	bsonDecodeContext bsoncodec.DecodeContext
}

// NewRecordsService
func NewRecordsService() RecordsRepository {
	recordsCollection := db.MongoClient.Collection("records")
	return &recordsRepository{
		recordsCollection: recordsCollection,
		bsonDecodeContext: bsoncodec.DecodeContext{
			Registry: bson.DefaultRegistry,
			Truncate: true,
		},
	}
}

// Get records from db between startDate, endDate and having totalCount between minCount, maxCount
func (rr *recordsRepository) GetRecords(
	startDate time.Time, endDate time.Time, minCount int, maxCount int,
) ([]models.Record, error) {
	cursor, err := rr.recordsCollection.Aggregate(context.Background(), []bson.M{
		{
			"$match": bson.M{
				"createdAt": bson.M{
					"$gte": startDate,
					"$lte": endDate,
				},
			},
		},
		{
			"$project": bson.M{
				"_id":        0,
				"key":        1,
				"createdAt":  1,
				"totalCount": bson.M{"$sum": "$counts"},
			},
		},
		{
			"$match": bson.M{
				"totalCount": bson.M{
					"$gte": minCount,
					"$lte": maxCount,
				},
			},
		},
	})

	var records []models.Record
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		record := &models.Record{}
		err := bson.UnmarshalWithContext(rr.bsonDecodeContext, cursor.Current, record)
		if err != nil {
			return nil, err
		}
		records = append(records, *record)
	}
	return records, err
}
