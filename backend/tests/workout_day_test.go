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
		Name:     "Test User",
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
func TestCreateWorkoutDay_Integration(t *testing.T) {
	router, db, userID, flowID := setupWorkoutDayTestEnvironment()
	defer db.Migrator().DropTable(&models.WorkoutDay{}, &models.User{}, &models.Flow{})

	t.Run("Success with valid flow", func(t *testing.T) {
		requestBody := bytes.NewBufferString(`{
			"title": "Leg Day",
			"day": "Monday",
			"duration": "60 minutes"
		}`)

		req, _ := http.NewRequest("POST", "/flows/"+flowID.String()+"/workout-days", requestBody)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)

		var response models.WorkoutDay
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Leg Day", response.Title)

		var dbWorkoutDay models.WorkoutDay
		db.Preload("User").Preload("Flow").First(&dbWorkoutDay, response.ID)
		assert.Equal(t, userID, dbWorkoutDay.UserID)
		assert.Equal(t, flowID, dbWorkoutDay.FlowID)
	})

	t.Run("Invalid input returns 400", func(t *testing.T) {
		requestBody := bytes.NewBufferString(`{
			"title": "",
			"day": "Monday"
		}`)

		req, _ := http.NewRequest("POST", "/flows/"+flowID.String()+"/workout-days", requestBody)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
func TestGetWorkoutDay_Integration(t *testing.T) {
	router, db, _, flowID := setupWorkoutDayTestEnvironment()
	defer db.Migrator().DropTable(&models.WorkoutDay{}, &models.User{}, &models.Flow{})

	workoutDay := models.WorkoutDay{
		ID:       uuid.New(),
		Title:    "Test Workout",
		Day:      "Tuesday",
		Duration: "45 minutes",
		UserID:   getTestUserID(db),
		FlowID:   flowID,
	}
	db.Create(&workoutDay)

	t.Run("Success", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/workout-days/"+workoutDay.ID.String(), nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response models.WorkoutDay
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, workoutDay.ID, response.ID)
	})

}
