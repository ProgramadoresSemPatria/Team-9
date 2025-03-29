package handlers

import (
	"net/http"
	"time"

	"github.com/ProgramadoresSemPatria/Team-9/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateFlow(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var input models.FlowInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	flow := models.Flow{
		Title:  input.Title,
		Level:  input.Level,
		Cover:  input.Cover,
		UserID: uuid.MustParse(userID.(string)),
	}

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Create(&flow).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create flow"})
		return
	}

	db.Preload("User").First(&flow, flow.ID)

	c.JSON(http.StatusCreated, flow)
}

func GetUserFlows(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	var flows []models.Flow

	if err := db.Preload("User").Where("user_id = ?", userID).Find(&flows).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch flows"})
		return
	}

	c.JSON(http.StatusOK, flows)
}

func GetFlow(c *gin.Context) {
	id := c.Param("id")
	db := c.MustGet("db").(*gorm.DB)

	var flow models.Flow
	if err := db.Preload("User").First(&flow, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Flow not found"})
		return
	}

	c.JSON(http.StatusOK, flow)
}

func UpdateFlow(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	id := c.Param("id")
	db := c.MustGet("db").(*gorm.DB)

	var existingFlow models.Flow
	if err := db.First(&existingFlow, "id = ? AND user_id = ?", id, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Flow not found or not owned by user"})
		return
	}

	var input models.FlowUpdate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := make(map[string]interface{})
	updates["updated_at"] = time.Now()

	if input.Title != "" {
		updates["title"] = input.Title
	}
	if input.Level != "" {
		updates["level"] = input.Level
	}
	if input.Cover != "" {
		updates["cover"] = input.Cover
	}

	if err := db.Model(&existingFlow).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update flow"})
		return
	}

	db.Preload("User").First(&existingFlow, existingFlow.ID)
	c.JSON(http.StatusOK, existingFlow)
}

func DeleteFlow(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	id := c.Param("id")
	db := c.MustGet("db").(*gorm.DB)

	var flow models.Flow
	if err := db.First(&flow, "id = ? AND user_id = ?", id, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Flow not found or not owned by user"})
		return
	}

	if err := db.Delete(&flow).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete flow"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Flow deleted successfully"})
}
