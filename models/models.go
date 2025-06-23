package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"
)

// StringArray is a custom type for []string to handle JSON in MySQL
// Implements Scanner and Valuer interfaces
// for GORM to store as JSON

type StringArray []string

func (a *StringArray) Scan(value interface{}) error {
	if value == nil {
		*a = []string{}
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, a)
}

func (a StringArray) Value() (driver.Value, error) {
	return json.Marshal(a)
}

type State struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Name        string         `gorm:"not null" json:"name"`
	Slug        string         `gorm:"uniqueIndex;not null" json:"slug"`
	PhoneNumber string         `gorm:"not null" json:"phone_number"`
	Heading     string         `json:"heading"`
	SubHeading  string         `json:"sub_heading"`
	Content     string         `json:"content"`
	Models      []Model        `gorm:"foreignKey:StateID" json:"models,omitempty"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

type Model struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	StateID     uint           `gorm:"not null" json:"state_id"`
	PhoneNumber string         `gorm:"not null" json:"phone_number"`
	Description string         `json:"description"`
	Name        string         `json:"name"`
	Heading     string         `json:"heading"`
	ProfileImg  string         `json:"profile_img"`
	BannerImg   string         `json:"banner_img"`
	Services    StringArray    `gorm:"type:json" json:"services"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

type FAQ struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GlobalPhone struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	PhoneNumber string    `gorm:"not null" json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
