package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Flow struct {
	ID        uuid.UUID `json:"id,omitempty" gorm:"type:uuid;primary_key;"`
	Title     string    `json:"title"`
	Level     string    `json:"level"`
	Cover     string    `json:"cover"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid"`
	User      User      `json:"user" gorm:"foreignKey:UserID;references:ID"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type FlowInput struct {
	Title string `json:"title" binding:"required"`
	Level string `json:"level" binding:"required"`
	Cover string `json:"cover"`
}

func (f *Flow) BeforeCreate(d *gorm.DB) (err error) {
	f.ID = uuid.New()
	return
}
