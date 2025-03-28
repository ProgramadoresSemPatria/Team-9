func CreateWorkoutDay(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	flowID := c.Param("flowId")
	var input models.WorkouDayInput

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
