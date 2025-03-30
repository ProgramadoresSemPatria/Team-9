func CreateExercise(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	userID := c.MustGet("userID").(string)

	workoutDayID := c.Param("id")
	if workoutDayID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Workout day ID is required"})
		return
	}

	var input models.ExerciseInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var workoutDay models.WorkoutDay
	if err := db.Where("id = ? AND user_id = ?", workoutDayID, userID).First(&workoutDay).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Workout day not found"})
		return
	}

	exercise := models.Exercise{
		Title:        input.Title,
		MuscleGroup:  input.MuscleGroup,
		Repetitions:  input.Repetitions,
		Sets:         input.Sets,
		WorkoutDayID: workoutDay.ID,
		UserID:       uuid.MustParse(userID),
	}

	if err := db.Create(&exercise).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create exercise"})
		return
	}

	c.JSON(http.StatusCreated, exercise)
}

func GetExercise(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	userID := c.MustGet("userID").(string)

	exerciseID := c.Param("id")
	if exerciseID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Exercise ID is required"})
		return
	}

	var exercise models.Exercise
	if err := db.Where("id = ? AND user_id = ?", exerciseID, userID).First(&exercise).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Exercise not found"})
		return
	}

	c.JSON(http.StatusOK, exercise)
}

func GetExercisesByWorkoutDay(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	userID := c.MustGet("userID").(string)

	workoutDayID := c.Param("id")
	if workoutDayID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Workout day ID is required"})
		return
	}

	var workoutDay models.WorkoutDay
	if err := db.Where("id = ? AND user_id = ?", workoutDayID, userID).First(&workoutDay).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Workout day not found"})
		return
	}

	var exercises []models.Exercise
	if err := db.Where("workout_day_id = ?", workoutDayID).Find(&exercises).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch exercises"})
		return
	}

	c.JSON(http.StatusOK, exercises)
}

func UpdateExercise(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	userID := c.MustGet("userID").(string)

	exerciseID := c.Param("id")
	if exerciseID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Exercise ID is required"})
		return
	}

	var input models.ExerciseuUpdateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var exercise models.Exercise
	if err := db.Where("id = ? AND user_id = ?", exerciseID, userID).First(&exercise).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Exercise not found"})
		return
	}

	if input.Title != "" {
		exercise.Title = input.Title
	}
	if input.MuscleGroup != "" {
		exercise.MuscleGroup = input.MuscleGroup
	}
	if input.Repetitions != 0 {
		exercise.Repetitions = input.Repetitions
	}
	if input.Sets != 0 {
		exercise.Sets = input.Sets
	}

	if err := db.Save(&exercise).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update exercise"})
		return
	}

	c.JSON(http.StatusOK, exercise)
}

func DeleteExercise(c *gin.Context) {
