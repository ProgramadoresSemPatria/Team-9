package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Flow struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
	Title     string    `json:"title" gorm:"size:255;not null"`
	Level     string    `json:"level" gorm:"size:50;not null"`
	Cover     string    `json:"cover" gorm:"size:255"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null;index"`
	User      User      `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
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
