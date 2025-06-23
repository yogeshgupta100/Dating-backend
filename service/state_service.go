package service

import (
	"model/models"
	"model/repository"
)

type StateService struct {
	stateRepo *repository.StateRepository
	modelRepo *repository.ModelRepository
}

func NewStateService(stateRepo *repository.StateRepository, modelRepo *repository.ModelRepository) *StateService {
	return &StateService{
		stateRepo: stateRepo,
		modelRepo: modelRepo,
	}
}

func (s *StateService) CreateState(state *models.State) error {
	return s.stateRepo.Create(state)
}

func (s *StateService) GetAllStates() ([]models.State, error) {
	return s.stateRepo.GetAll()
}

func (s *StateService) GetStateByID(id uint) (*models.State, error) {
	return s.stateRepo.GetByID(id)
}

func (s *StateService) GetStateBySlug(slug string) (*models.State, error) {
	return s.stateRepo.GetBySlug(slug)
}

func (s *StateService) GetModelsByStateID(stateID uint) ([]models.Model, error) {
	return s.modelRepo.GetByStateID(stateID)
}

func (s *StateService) UpdateState(state *models.State) error {
	return s.stateRepo.Update(state)
}

func (s *StateService) DeleteState(id uint) error {
	return s.stateRepo.Delete(id)
}
