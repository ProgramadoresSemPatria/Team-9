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

	// Match app routes
	router.POST("/flows/:id/workout-days", handlers.CreateWorkoutDay)
	router.GET("/workout-days/:id", handlers.GetWorkoutDay)
	router.GET("/flows/:id/workout-days", handlers.GetWorkoutDaysByFlow)
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
			"duration": 60
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
		Duration: 45,
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
		nonexistentID := uuid.New()
		req, _ := http.NewRequest("GET", "/workout-days/"+nonexistentID.String(), nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

func TestGetWorkoutDaysByFlow_Integration(t *testing.T) {
	router, db, _, flowID := setupWorkoutDayTestEnvironment()
	defer db.Migrator().DropTable(&models.WorkoutDay{}, &models.User{}, &models.Flow{})

	workoutDays := []models.WorkoutDay{
		{ID: uuid.New(), Title: "Day 1", Day: "Monday", Duration: 60, UserID: getTestUserID(db), FlowID: flowID},
		{ID: uuid.New(), Title: "Day 2", Day: "Wednesday", Duration: 45, UserID: getTestUserID(db), FlowID: flowID},
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
		newFlowID := uuid.New()
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
		ID:       uuid.New(),
		Title:    "Original Title",
		Day:      "Friday",
		Duration: 30,
		UserID:   getTestUserID(db),
		FlowID:   flowID,
	}
	db.Create(&workoutDay)

	t.Run("Success with valid input", func(t *testing.T) {
		requestBody := bytes.NewBufferString(`{
			"title": "Updated Title",
			"day": "Friday",
			"duration": 45
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
		assert.EqualValues(t, 45, response.Duration)

		var dbWorkoutDay models.WorkoutDay
		db.First(&dbWorkoutDay, workoutDay.ID)
		assert.Equal(t, "Updated Title", dbWorkoutDay.Title)
	})

	t.Run("Invalid input returns 400", func(t *testing.T) {
		requestBody := bytes.NewBufferString(`{
			"title": "",
			"duration": 45
		}`)

		req, _ := http.NewRequest("PUT", "/workout-days/"+workoutDay.ID.String(), requestBody)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Nonexistent workout day returns 404", func(t *testing.T) {
		nonexistentID := uuid.New()
		requestBody := bytes.NewBufferString(`{
			"title": "New Title",
			"day": "Monday",
			"duration": 60
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
			ID:       uuid.New(),
			Title:    "To Delete",
			Day:      "Sunday",
			Duration: 60,
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
		nonexistentID := uuid.New()
		req, _ := http.NewRequest("DELETE", "/workout-days/"+nonexistentID.String(), nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

func getTestUserID(db *gorm.DB) uuid.UUID {
	var user models.User
	if err := db.Where("email = ?", "test@example.com").First(&user).Error; err != nil {
		panic("Test user not found")
	}
	return user.ID
}
