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

