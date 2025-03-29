package migrations

import (
	"github.com/ProgramadoresSemPatria/Team-9/internal/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	createTables(db)
}

func createTables(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate((&models.Flow{}))
	db.AutoMigrate((&models.WorkoutDay{}))
}
