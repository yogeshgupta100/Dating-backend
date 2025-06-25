package service

import (
	"model/models"
	"model/repository"
)

type ModelService struct {
	modelRepo *repository.ModelRepository
}

func NewModelService(modelRepo *repository.ModelRepository) *ModelService {
	return &ModelService{
		modelRepo: modelRepo,
	}
}

func (s *ModelService) CreateModel(model *models.Model) error {
	return s.modelRepo.Create(model)
}

func (s *ModelService) GetAllModels() ([]models.Model, error) {
	return s.modelRepo.GetAll()
}

func (s *ModelService) GetModelByID(id uint) (*models.Model, error) {
	return s.modelRepo.GetByID(id)
}

func (s *ModelService) GetModelsByStateID(stateID uint) ([]models.Model, error) {
	return s.modelRepo.GetByStateID(stateID)
}

func (s *ModelService) UpdateModel(model *models.Model) error {
	return s.modelRepo.Update(model)
}

func (s *ModelService) DeleteModel(id uint) error {
	return s.modelRepo.Delete(id)
}

func (s *ModelService) GetModelsByHeading(heading string) ([]models.Model, error) {
	return s.modelRepo.GetByHeading(heading)
}

func (s *ModelService) GetModelsBySlug(slug string) ([]models.Model, error) {
	return s.modelRepo.GetBySlug(slug)
}
