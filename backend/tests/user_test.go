func setupRouter() (*gin.Engine, *gorm.DB) {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&models.User{})

	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	router.POST("/create", handlers.CreateUserHandler)
	router.POST("/login", handlers.LoginHandler)
	router.GET("/profile", handlers.AuthMiddleware(), handlers.ProfileHandler)
	return router, db
}
func TestCreateUserHandler(t *testing.T) {
	router, db := setupRouter()
	defer db.Exec("DROP TABLE users")

	userInput := models.SignInInput{Name: "Teste" , Email: "test@example.com", Password: "password123"}
	body, _ := json.Marshal(userInput)

	req, _ := http.NewRequest("POST", "/create", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var user models.User
	db.Where("email = ?", userInput.Email).First(&user)
	assert.NotEmpty(t, user.ID)
	assert.True(t, user.Verified)
}
