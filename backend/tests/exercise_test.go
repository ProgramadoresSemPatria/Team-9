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

func setupExerciseTestEnvironment() (*gin.Engine, *gorm.DB, uuid.UUID, uuid.UUID) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{}, &models.WorkoutDay{}, &models.Exercise{})

	testUser := models.User{
		ID:       uuid.New(),
		Email:    "exercise@example.com",
		Password: "secure_password",
	}
	db.Create(&testUser)

	workoutDay := models.WorkoutDay{
		ID:     uuid.New(),
		Title:  "Test Day",
		UserID: testUser.ID,
	}
	db.Create(&workoutDay)

	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Set("userID", testUser.ID.String())
	})

	router.POST("/workout-days/:id/exercises", handlers.CreateExercise)

	return router, db, testUser.ID, workoutDay.ID
}

func TestCreateExercise_Integration(t *testing.T) {
	router, db, _, workoutDayID := setupExerciseTestEnvironment()
	defer db.Migrator().DropTable(&models.Exercise{}, &models.WorkoutDay{}, &models.User{})

	t.Run("Success", func(t *testing.T) {
		body := bytes.NewBufferString(`{
			"title": "Push Up",
			"muscle_group": "Chest",
			"repetitions": 15,
			"sets": 3
		}`)

		req, _ := http.NewRequest("POST", "/workout-days/"+workoutDayID.String()+"/exercises", body)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)

		var response models.Exercise
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Push Up", response.Title)

		var dbExercise models.Exercise
		result := db.First(&dbExercise, "id = ?", response.ID)
		assert.NoError(t, result.Error)
		assert.Equal(t, response.ID, dbExercise.ID)
	})

	t.Run("Missing Fields", func(t *testing.T) {
		body := bytes.NewBufferString(`{"title": ""}`)

		req, _ := http.NewRequest("POST", "/workout-days/"+workoutDayID.String()+"/exercises", body)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Workout Day Not Found", func(t *testing.T) {
		fakeWorkoutID := uuid.New().String()
		body := bytes.NewBufferString(`{
			"title": "Pull Up",
			"muscle_group": "Back",
			"repetitions": 10,
			"sets": 4
		}`)

		req, _ := http.NewRequest("POST", "/workout-days/"+fakeWorkoutID+"/exercises", body)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

func TestGetExercise_Integration(t *testing.T) {
	router, db, userID, workoutDayID := setupExerciseTestEnvironment()
	defer db.Migrator().DropTable(&models.Exercise{}, &models.WorkoutDay{}, &models.User{})

	exercise := models.Exercise{
		ID:           uuid.New(),
		Title:        "Deadlift",
		MuscleGroup:  "Back",
		Repetitions:  5,
		Sets:         5,
		WorkoutDayID: workoutDayID,
		UserID:       userID,
	}
	db.Create(&exercise)

	router.GET("/exercises/:id", handlers.GetExercise)

	t.Run("Success", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/exercises/"+exercise.ID.String(), nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response models.Exercise
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Deadlift", response.Title)
	})

	t.Run("Not Found", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/exercises/"+uuid.New().String(), nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

func TestGetExercisesByWorkoutDay_Integration(t *testing.T) {
	router, db, userID, workoutDayID := setupExerciseTestEnvironment()
	defer db.Migrator().DropTable(&models.Exercise{}, &models.WorkoutDay{}, &models.User{})

	router.GET("/workout-days/:id/exercises", handlers.GetExercisesByWorkoutDay)

	exercises := []models.Exercise{
		{ID: uuid.New(), Title: "Squat", MuscleGroup: "Legs", Repetitions: 10, Sets: 4, WorkoutDayID: workoutDayID, UserID: userID},
		{ID: uuid.New(), Title: "Lunge", MuscleGroup: "Legs", Repetitions: 12, Sets: 3, WorkoutDayID: workoutDayID, UserID: userID},
	}
	for _, e := range exercises {
		db.Create(&e)
	}

	t.Run("Success", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/workout-days/"+workoutDayID.String()+"/exercises", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response []models.Exercise
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Len(t, response, 2)
	})

	t.Run("Workout Day Not Found", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/workout-days/"+uuid.New().String()+"/exercises", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

func TestUpdateExercise_Integration(t *testing.T) {
	router, db, userID, workoutDayID := setupExerciseTestEnvironment()
	defer db.Migrator().DropTable(&models.Exercise{}, &models.WorkoutDay{}, &models.User{})

	exercise := models.Exercise{
		ID:           uuid.New(),
		Title:        "Pull Up",
		MuscleGroup:  "Back",
		Repetitions:  10,
		Sets:         3,
		WorkoutDayID: workoutDayID,
		UserID:       userID,
	}
	db.Create(&exercise)

	router.PUT("/exercises/:id", handlers.UpdateExercise)

	t.Run("Success", func(t *testing.T) {
		body := bytes.NewBufferString(`{
			"title": "Wide Pull Up",
			"repetitions": 12
		}`)

		req, _ := http.NewRequest("PUT", "/exercises/"+exercise.ID.String(), body)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response models.Exercise
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Wide Pull Up", response.Title)
		assert.Equal(t, int64(12), response.Repetitions)
	})

	t.Run("Not Found", func(t *testing.T) {
		body := bytes.NewBufferString(`{"title": "Ghost Update"}`)
		req, _ := http.NewRequest("PUT", "/exercises/"+uuid.New().String(), body)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("Invalid Input", func(t *testing.T) {
		body := bytes.NewBufferString(`{"title": 123}`)

		req, _ := http.NewRequest("PUT", "/exercises/"+exercise.ID.String(), body)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestDeleteExercise_Integration(t *testing.T) {
	router, db, userID, workoutDayID := setupExerciseTestEnvironment()
	defer db.Migrator().DropTable(&models.Exercise{}, &models.WorkoutDay{}, &models.User{})

	router.DELETE("/exercises/:id", handlers.DeleteExercise)

	t.Run("Success", func(t *testing.T) {
		exercise := models.Exercise{
			ID:           uuid.New(),
			Title:        "Burpee",
			MuscleGroup:  "Full Body",
			Repetitions:  20,
			Sets:         4,
			WorkoutDayID: workoutDayID,
			UserID:       userID,
		}
		db.Create(&exercise)

		req, _ := http.NewRequest("DELETE", "/exercises/"+exercise.ID.String(), nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var dbExercise models.Exercise
		result := db.First(&dbExercise, exercise.ID)
		assert.Error(t, result.Error)
		assert.Equal(t, gorm.ErrRecordNotFound, result.Error)
	})

	t.Run("Not Found", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/exercises/"+uuid.New().String(), nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}
