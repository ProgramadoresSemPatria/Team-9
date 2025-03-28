func setupWorkoutDayTestEnvironment() (*gin.Engine, *gorm.DB, uuid.UUID, uuid.UUID) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.WorkoutDay{}, &models.User{}, &models.Flow{})

	testUser := models.User{
		ID:       uuid.New(),
		Email:    "test@example.com",
		Password: "hashed_password",
		Name:     "test",
	}
	db.Create(&testUser)

	testFlow := models.Flow{
		ID:     uuid.New(),
		Title:  "Test Flow",
		UserID: testUser.ID,
	}
	db.Create(&testFlow)

	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Set("userID", testUser.ID.String())
	})

	router.POST("/flows/:flowId/workout-days", handlers.CreateWorkoutDay)
	router.GET("/workout-days/:id", handlers.GetWorkoutDay)
	router.GET("/flows/:flowId/workout-days", handlers.GetWorkoutDaysByFlow)
	router.PUT("/workout-days/:id", handlers.UpdateWorkoutDay)
	router.DELETE("/workout-days/:id", handlers.DeleteWorkoutDay)

	return router, db, testUser.ID, testFlow.ID
}
