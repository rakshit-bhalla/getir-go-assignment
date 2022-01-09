package repositories

import (
	"rakshit.dev/m/src/db"
	"rakshit.dev/m/src/models"
)

// InMemoryRepository
type InMemoryRepository interface {
	Write(keyValue models.KeyValue) (*models.KeyValue, error)
	Read(key string) (*models.KeyValue, error)
}

type inMemoryRepository struct {
	datastore map[string]string
}

// NewInMemoryService
func NewInMemoryService() InMemoryRepository {
	return &inMemoryRepository{
		datastore: db.Datastore,
	}
}

// Writes value for a key into in-memory db
func (ims *inMemoryRepository) Write(keyValue models.KeyValue) (*models.KeyValue, error) {
	ims.datastore[keyValue.Key] = keyValue.Value
	return &keyValue, nil
}

// Reads value for a key from in-memory db
func (ims *inMemoryRepository) Read(key string) (*models.KeyValue, error) {
	value := ims.datastore[key]
	if value == "" {
		return nil, nil
	}
	return &models.KeyValue{Key: key, Value: ims.datastore[key]}, nil
}
