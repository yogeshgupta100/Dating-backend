package service

import (
	"model/models"
	"model/repository"
)

type FAQService struct {
	faqRepo *repository.FAQRepository
}

func NewFAQService(faqRepo *repository.FAQRepository) *FAQService {
	return &FAQService{faqRepo: faqRepo}
}

func (s *FAQService) GetFAQ() (*models.FAQ, error) {
	return s.faqRepo.GetFAQ()
}

func (s *FAQService) CreateOrUpdateFAQ(name string) (*models.FAQ, error) {
	return s.faqRepo.CreateOrUpdateFAQ(name)
}
