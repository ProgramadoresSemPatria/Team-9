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

