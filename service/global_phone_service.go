package service

import (
	"model/models"
	"model/repository"
)

type GlobalPhoneService struct {
	globalPhoneRepo *repository.GlobalPhoneRepository
}

func NewGlobalPhoneService(globalPhoneRepo *repository.GlobalPhoneRepository) *GlobalPhoneService {
	return &GlobalPhoneService{
		globalPhoneRepo: globalPhoneRepo,
	}
}

func (s *GlobalPhoneService) CreateOrUpdateGlobalPhone(globalPhone *models.GlobalPhone) error {
	return s.globalPhoneRepo.CreateOrUpdate(globalPhone)
}

func (s *GlobalPhoneService) GetGlobalPhone() (*models.GlobalPhone, error) {
	return s.globalPhoneRepo.Get()
}
