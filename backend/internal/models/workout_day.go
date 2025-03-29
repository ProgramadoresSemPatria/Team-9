package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WorkoutDay struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Title     string    `json:"title" gorm:"size:255;not null"`
	Day       string    `json:"day" gorm:"size:50;not null"`
	Duration  string    `json:"duration" gorm:"size:50;not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null;index"`
	User      User      `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	FlowID    uuid.UUID `json:"flow_id" gorm:"type:uuid;not null;index"`
	Flow      Flow      `json:"flow" gorm:"foreignKey:FlowID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type WorkouDayInput struct {
	Title    string `json:"title" binding:"required,min=3,max=255"`
	Day      string `json:"day" binding:"required,min=3,max=50"`
	Duration string `json:"duration" binding:"required,min=3,max=50"`
}

func (w *WorkoutDay) BeforeCreate(d *gorm.DB) (err error) {
	w.ID = uuid.New()
	return
}
