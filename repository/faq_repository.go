package repository

import (
	"model/models"

	"gorm.io/gorm"
)

type FAQRepository struct {
	db *gorm.DB
}

func NewFAQRepository(db *gorm.DB) *FAQRepository {
	return &FAQRepository{db: db}
}

func (r *FAQRepository) GetFAQ() (*models.FAQ, error) {
	var faq models.FAQ
	err := r.db.First(&faq).Error
	return &faq, err
}

func (r *FAQRepository) CreateOrUpdateFAQ(name string) (*models.FAQ, error) {
	var faq models.FAQ
	err := r.db.First(&faq).Error
	if err == gorm.ErrRecordNotFound {
		faq = models.FAQ{Name: name}
		err = r.db.Create(&faq).Error
		return &faq, err
	} else if err != nil {
		return nil, err
	}
	faq.Name = name
	err = r.db.Save(&faq).Error
	return &faq, err
}
