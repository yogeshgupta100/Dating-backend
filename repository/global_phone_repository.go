package repository

import (
	"model/models"

	"gorm.io/gorm"
)

type GlobalPhoneRepository struct {
	db *gorm.DB
}

func NewGlobalPhoneRepository(db *gorm.DB) *GlobalPhoneRepository {
	return &GlobalPhoneRepository{db: db}
}

func (r *GlobalPhoneRepository) CreateOrUpdate(globalPhone *models.GlobalPhone) error {
	// Check if any record exists
	var count int64
	r.db.Model(&models.GlobalPhone{}).Count(&count)

	if count == 0 {
		// No record exists, create new one
		return r.db.Create(globalPhone).Error
	} else {
		// Record exists, update the first one
		var existingGlobalPhone models.GlobalPhone
		if err := r.db.First(&existingGlobalPhone).Error; err != nil {
			return err
		}
		// Update only the phone_number field to preserve timestamps
		return r.db.Model(&existingGlobalPhone).Update("phone_number", globalPhone.PhoneNumber).Error
	}
}

func (r *GlobalPhoneRepository) Get() (*models.GlobalPhone, error) {
	var globalPhone models.GlobalPhone
	err := r.db.First(&globalPhone).Error
	if err != nil {
		return nil, err
	}
	return &globalPhone, nil
}
