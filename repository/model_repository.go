package repository

import (
	"model/models"

	"gorm.io/gorm"
)

type ModelRepository struct {
	db *gorm.DB
}

func NewModelRepository(db *gorm.DB) *ModelRepository {
	return &ModelRepository{db: db}
}

func (r *ModelRepository) Create(model *models.Model) error {
	return r.db.Create(model).Error
}

func (r *ModelRepository) GetAll() ([]models.Model, error) {
	var models []models.Model
	err := r.db.Find(&models).Error
	return models, err
}

func (r *ModelRepository) GetByID(id uint) (*models.Model, error) {
	var model models.Model
	err := r.db.First(&model, id).Error
	return &model, err
}

func (r *ModelRepository) GetByStateID(stateID uint) ([]models.Model, error) {
	var models []models.Model
	err := r.db.Where("state_id = ?", stateID).Find(&models).Error
	return models, err
}

func (r *ModelRepository) Update(model *models.Model) error {
	return r.db.Save(model).Error
}

func (r *ModelRepository) Delete(id uint) error {
	return r.db.Delete(&models.Model{}, id).Error
}

func (r *ModelRepository) GetByHeading(heading string) ([]models.Model, error) {
	var models []models.Model
	err := r.db.Where("heading = ?", heading).Find(&models).Error
	return models, err
}

func (r *ModelRepository) GetBySlug(slug string) ([]models.Model, error) {
	var models []models.Model
	err := r.db.Where("slug = ?", slug).Find(&models).Error
	return models, err
}

func (r *ModelRepository) DeleteByStateID(stateID uint) error {
	return r.db.Where("state_id = ?", stateID).Delete(&models.Model{}).Error
}

func (r *ModelRepository) CountByStateID(stateID uint) (int64, error) {
	var count int64
	err := r.db.Model(&models.Model{}).Where("state_id = ?", stateID).Count(&count).Error
	return count, err
}
