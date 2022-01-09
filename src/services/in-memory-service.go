package services

import (
	"rakshit.dev/m/src/models"
	"rakshit.dev/m/src/repositories"
)

// InMemoryService
type InMemoryService interface {
	Write(keyValue models.KeyValue) (*models.KeyValue, error)
	Read(key string) (*models.KeyValue, error)
}

type inMemoryService struct {
	inMemoryRepository repositories.InMemoryRepository
}

// NewInMemoryService
func NewInMemoryService(
	inMemoryRepository repositories.InMemoryRepository,
) InMemoryService {
	return &inMemoryService{
		inMemoryRepository: inMemoryRepository,
	}
}

// Writes value for a key into in-memory db
func (ims *inMemoryService) Write(keyValue models.KeyValue) (*models.KeyValue, error) {
	return ims.inMemoryRepository.Write(keyValue)
}

// Reads value for a key from in-memory db
func (ims *inMemoryService) Read(key string) (*models.KeyValue, error) {
	return ims.inMemoryRepository.Read(key)
}
