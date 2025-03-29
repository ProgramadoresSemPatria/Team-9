package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ProgramadoresSemPatria/Team-9/internal/handlers"
	"github.com/ProgramadoresSemPatria/Team-9/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupWorkoutDayTestEnvironment() (*gin.Engine, *gorm.DB, uuid.UUID, uuid.UUID) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.WorkoutDay{}, &models.User{}, &models.Flow{})

	testUser := models.User{
		ID:       uuid.MustParse("ff44a5c3-9c9f-4adc-b26b-d51f569af8ea"),
		Email:    "test@example.com",
		Password: "hashed_password",
		Name:     "Test User",
	}
	db.Create(&testUser)

	testFlow := models.Flow{
		ID:     uuid.MustParse("42d795f3-c9ac-46d9-8356-97af5a2616e7"),
		Title:  "Test Flow",
		UserID: testUser.ID,
	}
	db.Create(&testFlow)

	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Set("userID", testUser.ID.String())
	})

	// Mock auth middleware for testing
	authGroup := router.Group("/")
	authGroup.Use(func(c *gin.Context) {
		c.Set("userID", testUser.ID.String())
		c.Next()
	})

	// Set up routes to match your application
	authGroup.POST("/flows/:id/workout-days", handlers.CreateWorkoutDay)
	authGroup.GET("/flows/:id/workout-days", handlers.GetWorkoutDaysByFlow)
	authGroup.GET("/workout-days/:id", handlers.GetWorkoutDay)
	authGroup.PUT("/workout-days/:id", handlers.UpdateWorkoutDay)
	authGroup.DELETE("/workout-days/:id", handlers.DeleteWorkoutDay)

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

		assert.Equal(t, http.StatusCreated, w.Code, "Expected status code 201")

		var response models.WorkoutDay
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Leg Day", response.Title)

		var dbWorkoutDay models.WorkoutDay
		result := db.Preload("User").Preload("Flow").First(&dbWorkoutDay, "id = ?", response.ID)
		assert.NoError(t, result.Error)
		assert.Equal(t, userID, dbWorkoutDay.UserID)
		assert.Equal(t, flowID, dbWorkoutDay.FlowID)
	})

}

func TestGetWorkoutDay_Integration(t *testing.T) {
	router, db, _, flowID := setupWorkoutDayTestEnvironment()
	defer db.Migrator().DropTable(&models.WorkoutDay{}, &models.User{}, &models.Flow{})

	workoutDay := models.WorkoutDay{
		ID:       uuid.MustParse("e14eb1b0-85fa-4392-a67a-3f34a068d8d6"),
		Title:    "Test Workout",
		Day:      "Tuesday",
		Duration: "45 minutes",
		UserID:   getTestUserID(db),
		FlowID:   flowID,
	}
	db.Create(&workoutDay)

	t.Run("Success with existing workout day", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/workout-days/"+workoutDay.ID.String(), nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response models.WorkoutDay
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, workoutDay.ID, response.ID)
		assert.Equal(t, "Test Workout", response.Title)
	})

	t.Run("Nonexistent workout day returns 404", func(t *testing.T) {
		nonexistentID := uuid.MustParse("74709a57-97c0-4264-8375-e9b14ad6c2c8")
		req, _ := http.NewRequest("GET", "/workout-days/"+nonexistentID.String(), nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

func TestGetWorkoutDaysByFlow_Integration(t *testing.T) {
	router, db, _, flowID := setupWorkoutDayTestEnvironment()
	defer db.Migrator().DropTable(&models.WorkoutDay{}, &models.User{}, &models.Flow{})

	// Create test workout days
	workoutDays := []models.WorkoutDay{
		{
			ID:       uuid.New(),
			Title:    "Day 1",
			Day:      "Monday",
			Duration: "60 mins",
			UserID:   getTestUserID(db),
			FlowID:   flowID,
		},
		{
			ID:       uuid.New(),
			Title:    "Day 2",
			Day:      "Wednesday",
			Duration: "45 mins",
			UserID:   getTestUserID(db),
			FlowID:   flowID,
		},
	}
	for _, wd := range workoutDays {
		db.Create(&wd)
	}

	t.Run("Success with existing flow", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/flows/"+flowID.String()+"/workout-days", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response []models.WorkoutDay
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Len(t, response, 2)
	})

	t.Run("Empty list for new flow returns 200", func(t *testing.T) {
		newFlowID := uuid.MustParse("e5770a66-537d-42e5-adcb-c1070b886a8d")
		req, _ := http.NewRequest("GET", "/flows/"+newFlowID.String()+"/workout-days", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response []models.WorkoutDay
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Empty(t, response)
	})
}

func TestUpdateWorkoutDay_Integration(t *testing.T) {
	router, db, _, flowID := setupWorkoutDayTestEnvironment()
	defer db.Migrator().DropTable(&models.WorkoutDay{}, &models.User{}, &models.Flow{})

	workoutDay := models.WorkoutDay{
		ID:       uuid.MustParse("710e6d49-48d6-4bdd-b690-e6f4b58da7c4"),
		Title:    "Original Title",
		Day:      "Friday",
		Duration: "30 minutes",
		UserID:   getTestUserID(db),
		FlowID:   flowID,
	}
	db.Create(&workoutDay)

	t.Run("Success with valid input", func(t *testing.T) {
		requestBody := bytes.NewBufferString(`{
			"title": "Updated Title",
			"day": "Friday",
			"duration": "45 minutes"
		}`)

		req, _ := http.NewRequest("PUT", "/workout-days/"+workoutDay.ID.String(), requestBody)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response models.WorkoutDay
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Updated Title", response.Title)
		assert.Equal(t, "45 minutes", response.Duration)

		var dbWorkoutDay models.WorkoutDay
		db.First(&dbWorkoutDay, workoutDay.ID)
		assert.Equal(t, "Updated Title", dbWorkoutDay.Title)
	})

	t.Run("Nonexistent workout day returns 404", func(t *testing.T) {
		nonexistentID := uuid.MustParse("3ab2a2e6-9787-459b-8a66-55232e1087ef")
		requestBody := bytes.NewBufferString(`{
			"title": "New Title",
			"day": "Monday",
			"duration": "60 minutes"
		}`)

		req, _ := http.NewRequest("PUT", "/workout-days/"+nonexistentID.String(), requestBody)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

func TestDeleteWorkoutDay_Integration(t *testing.T) {
	router, db, _, flowID := setupWorkoutDayTestEnvironment()
	defer db.Migrator().DropTable(&models.WorkoutDay{}, &models.User{}, &models.Flow{})

	t.Run("Success with existing workout day", func(t *testing.T) {
		workoutDay := models.WorkoutDay{
			ID:       uuid.MustParse("c2da191c-4c90-4947-b1f0-25d0dd1adfa1"),
			Title:    "To Delete",
			Day:      "Sunday",
			Duration: "60 minutes",
			UserID:   getTestUserID(db),
			FlowID:   flowID,
		}
		db.Create(&workoutDay)

		req, _ := http.NewRequest("DELETE", "/workout-days/"+workoutDay.ID.String(), nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var dbWorkoutDay models.WorkoutDay
		result := db.First(&dbWorkoutDay, workoutDay.ID)
		assert.Error(t, result.Error)
		assert.Equal(t, gorm.ErrRecordNotFound, result.Error)
	})

	t.Run("Nonexistent workout day returns 404", func(t *testing.T) {
		nonexistentID := uuid.MustParse("8cd8e6ca-5ba1-48dc-9715-a08fe99d35ef")
		req, _ := http.NewRequest("DELETE", "/workout-days/"+nonexistentID.String(), nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

func getTestUserID(db *gorm.DB) uuid.UUID {
	var user models.User
	db.Where("email = ?", "test@example.com").First(&user)
	return user.ID
}
