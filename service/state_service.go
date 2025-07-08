package service

import (
	"fmt"
	"log"
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
	// First, check if the state exists
	state, err := s.stateRepo.GetByID(id)
	if err != nil {
		log.Printf("Error finding state ID %d: %v", id, err)
		return fmt.Errorf("state not found: %w", err)
	}

	log.Printf("Found state: %s (ID: %d), proceeding with deletion", state.Name, id)

	// Count models associated with this state
	modelCount, err := s.modelRepo.CountByStateID(id)
	if err != nil {
		log.Printf("Error counting models for state ID %d: %v", id, err)
		// Continue with deletion even if count fails
	} else {
		log.Printf("Found %d models associated with state ID %d", modelCount, id)
	}

	// Delete all models associated with this state
	log.Printf("Deleting all models for state ID: %d", id)
	if err := s.modelRepo.DeleteByStateID(id); err != nil {
		log.Printf("Error deleting models for state ID %d: %v", id, err)
		return fmt.Errorf("failed to delete models for state: %w", err)
	}

	// Then delete the state
	log.Printf("Deleting state ID: %d", id)
	if err := s.stateRepo.Delete(id); err != nil {
		log.Printf("Error deleting state ID %d: %v", id, err)
		return fmt.Errorf("failed to delete state: %w", err)
	}

	log.Printf("Successfully deleted state ID: %d and all its models", id)
	return nil
}
