package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"strings"
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

// GenerateSlug converts a heading to a URL-friendly slug
func GenerateSlug(heading string) string {
	// Convert to lowercase
	slug := strings.ToLower(heading)

	// Replace spaces with hyphens
	slug = strings.ReplaceAll(slug, " ", "-")

	// Remove any other special characters that might cause issues
	// Keep only alphanumeric characters and hyphens
	var result strings.Builder
	for _, char := range slug {
		if (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9') || char == '-' {
			result.WriteRune(char)
		}
	}

	// Remove multiple consecutive hyphens
	slug = result.String()
	for strings.Contains(slug, "--") {
		slug = strings.ReplaceAll(slug, "--", "-")
	}

	// Remove leading and trailing hyphens
	slug = strings.Trim(slug, "-")

	return slug
}

type State struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Name        string         `gorm:"not null" json:"name"`
	Slug        string         `gorm:"type:varchar(191);uniqueIndex;not null" json:"slug"`
	PhoneNumber string         `gorm:"not null" json:"phone_number"`
	Heading     string         `json:"heading"`
	SubHeading  string         `json:"sub_heading"`
	Content     string         `json:"content"`
	SEOTitle    string         `json:"seo_title"`
	SEODesc     string         `json:"seo_desc"`
	SEOKeyword  string         `json:"seo_keyword"`
	FAQ         string         `json:"faq"`
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
	Slug        string         `gorm:"type:varchar(191);uniqueIndex" json:"slug"`
	ProfileImg  string         `json:"profile_img"`
	BannerImg   string         `json:"banner_img"`
	Services    StringArray    `gorm:"type:json" json:"services"`
	SEOTitle    string         `json:"seo_title"`
	SEODesc     string         `json:"seo_desc"`
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
