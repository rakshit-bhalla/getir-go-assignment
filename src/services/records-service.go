package services

import (
	"time"

	"rakshit.dev/m/src/models"
	"rakshit.dev/m/src/repositories"
)

// RecordsService
type RecordsService interface {
	GetRecords(
		startDate time.Time, endDate time.Time, minCount int, maxCount int,
	) ([]models.Record, error)
}

type recordsService struct {
	recordsRepository repositories.RecordsRepository
}

// NewRecordsService
func NewRecordsService(
	recordsRepository repositories.RecordsRepository,
) RecordsService {
	return &recordsService{
		recordsRepository: recordsRepository,
	}
}

// Get records from db between startDate, endDate and having totalCount between minCount, maxCount
func (rs *recordsService) GetRecords(
	startDate time.Time, endDate time.Time, minCount int, maxCount int,
) ([]models.Record, error) {
	return rs.recordsRepository.GetRecords(startDate, endDate, minCount, maxCount)
}
