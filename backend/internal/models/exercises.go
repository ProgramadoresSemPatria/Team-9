package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Exercise struct {
	ID           uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;"`
	Title        string     `json:"title" gorm:"size:255;not null"`
	MuscleGroup  string     `json:"muscle_group" gorm:"size:255;not null"`
	Repetitions  int64      `json:"repetitions" gorm:"size:255;not null"`
	Sets         int64      `json:"sets" gorm:"size:255;not null"`
	WorkoutDayID WorkoutDay `json:"workout_day_id" gorm:"foreignKey:WorkoutDayID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt    time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
}

