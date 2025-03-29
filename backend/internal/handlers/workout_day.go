package handlers

import (
	"net/http"
	"time"

	"github.com/ProgramadoresSemPatria/Team-9/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateWorkoutDay(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var input models.WorkouDayInput
	flowID := c.Param("id")

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.MustGet("db").(*gorm.DB)

	workoutDay := models.WorkoutDay{
		Title:    input.Title,
		Day:      input.Day,
		Duration: input.Duration,
		UserID:   uuid.MustParse(userID.(string)),
		FlowID:   uuid.MustParse(flowID),
	}

	if err := db.Create(&workoutDay).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create workout day"})
		return
	}

	db.Preload("User").Preload("Flow").First(&workoutDay, workoutDay.ID)
	c.JSON(http.StatusCreated, workoutDay)
}

func GetWorkoutDay(c *gin.Context) {
	id := c.Param("id")
	db := c.MustGet("db").(*gorm.DB)

	var workoutDay models.WorkoutDay
	if err := db.Preload("User").Preload("Flow").First(&workoutDay, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Workout day not found"})
		return
	}

	c.JSON(http.StatusOK, workoutDay)
}

func GetWorkoutDaysByFlow(c *gin.Context) {
	flowID := c.Param("flowId")
	db := c.MustGet("db").(*gorm.DB)

	var workoutDays []models.WorkoutDay
	if err := db.Preload("User").Preload("Flow").Where("flow_id = ?", flowID).Find(&workoutDays).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch workout days"})
		return
	}

	c.JSON(http.StatusOK, workoutDays)
}

func UpdateWorkoutDay(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	id := c.Param("id")
	db := c.MustGet("db").(*gorm.DB)

	var existingWorkoutDay models.WorkoutDay
	if err := db.First(&existingWorkoutDay, "id = ? AND user_id = ?", id, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Workout day not found or not owned by user"})
		return
	}

	var input models.WorkouDayInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := map[string]interface{}{
		"title":      input.Title,
		"day":        input.Day,
		"duration":   input.Duration,
		"updated_at": time.Now(),
	}

	if err := db.Model(&existingWorkoutDay).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update workout day"})
		return
	}

	db.Preload("User").Preload("Flow").First(&existingWorkoutDay, existingWorkoutDay.ID)
	c.JSON(http.StatusOK, existingWorkoutDay)
}

func DeleteWorkoutDay(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	id := c.Param("id")
	db := c.MustGet("db").(*gorm.DB)

	var workoutDay models.WorkoutDay
	if err := db.First(&workoutDay, "id = ? AND user_id = ?", id, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Workout day not found or not owned by user"})
		return
	}

	if err := db.Delete(&workoutDay).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete workout day"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Workout day deleted successfully"})
}
