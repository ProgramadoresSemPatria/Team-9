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
	Repetitions  int64      `json:"repetitions" gorm:"not null"`
	Sets         int64      `json:"sets" gorm:"not null"`
	WorkoutDayID uuid.UUID  `json:"workout_day_id" gorm:"type:uuid;not null;index"`
	WorkoutDay   WorkoutDay `json:"workout_day" gorm:"foreignKey:WorkoutDayID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID       uuid.UUID  `json:"user_id" gorm:"type:uuid;not null;index"`
	User         User       `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt    time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
}

type ExerciseInput struct {
	Title       string `json:"title" binding:"required,min=3,max=255"`
	MuscleGroup string `json:"muscle_group" binding:"required,min=3,max=255"`
	Repetitions int64  `json:"repetitions" binding:"required"`
	Sets        int64  `json:"sets" binding:"required"`
}

type ExerciseuUpdateInput struct {
	Title       string `json:"title"`
	MuscleGroup string `json:"muscle_group"`
	Repetitions int64  `json:"repetitions"`
	Sets        int64  `json:"sets"`
}

func (e *Exercise) BeforeCreate(d *gorm.DB) (err error) {
	e.ID = uuid.New()
	return
}
