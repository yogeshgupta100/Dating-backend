package repository

import (
	"model/models"

	"gorm.io/gorm"
)

type StateRepository struct {
	db *gorm.DB
}

func NewStateRepository(db *gorm.DB) *StateRepository {
	return &StateRepository{db: db}
}

func (r *StateRepository) Create(state *models.State) error {
	return r.db.Create(state).Error
}

func (r *StateRepository) GetAll() ([]models.State, error) {
	var states []models.State
	err := r.db.Find(&states).Error
	return states, err
}

func (r *StateRepository) GetByID(id uint) (*models.State, error) {
	var state models.State
	err := r.db.First(&state, id).Error
	if err != nil {
		return nil, err
	}
	return &state, nil
}

func (r *StateRepository) GetBySlug(slug string) (*models.State, error) {
	var state models.State
	err := r.db.Where("slug = ?", slug).First(&state).Error
	if err != nil {
		return nil, err
	}
	return &state, nil
}

func (r *StateRepository) GetModelsByStateID(stateID uint) ([]models.Model, error) {
	var models []models.Model
	err := r.db.Where("state_id = ?", stateID).Find(&models).Error
	return models, err
}

func (r *StateRepository) Update(state *models.State) error {
	return r.db.Save(state).Error
}

func (r *StateRepository) Delete(id uint) error {
	return r.db.Delete(&models.State{}, id).Error
}
