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

func setupTestEnvironment() (*gin.Engine, *gorm.DB, uuid.UUID) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Flow{}, &models.User{})

	testUser := models.User{
		ID:       uuid.New(),
		Email:    "test@example.com",
		Password: "hashed_password",
	}
	db.Create(&testUser)

	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Set("userID", testUser.ID.String())
	})

	router.POST("/flows", handlers.CreateFlow)
	router.GET("/flows", handlers.GetUserFlows)
	router.GET("/flows/:id", handlers.GetFlow)
	router.PUT("/flows/:id", handlers.UpdateFlow)
	router.DELETE("/flows/:id", handlers.DeleteFlow)

	return router, db, testUser.ID
}

func TestCreateFlow_Integration(t *testing.T) {
	router, db, _ := setupTestEnvironment()
	defer db.Migrator().DropTable(&models.Flow{}, &models.User{})

	t.Run("Success", func(t *testing.T) {
		requestBody := bytes.NewBufferString(`{
			"title": "Integration Test Flow",
			"level": "intermediate",
			"cover": "cover.jpg"
		}`)

		req, _ := http.NewRequest("POST", "/flows", requestBody)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)

		var response models.Flow
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Integration Test Flow", response.Title)

		var dbFlow models.Flow
		db.Preload("User").First(&dbFlow, response.ID)
		assert.Equal(t, response.ID, dbFlow.ID)
		assert.Equal(t, "test@example.com", dbFlow.User.Email)
	})

	t.Run("Invalid Input", func(t *testing.T) {
		requestBody := bytes.NewBufferString(`{
			"invalid": "json"
		}`)

		req, _ := http.NewRequest("POST", "/flows", requestBody)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Unauthorized", func(t *testing.T) {
		unauthRouter := gin.Default()
		unauthRouter.Use(func(c *gin.Context) {
			c.Set("db", db)
		})
		unauthRouter.POST("/flows", handlers.CreateFlow)

		requestBody := bytes.NewBufferString(`{
			"title": "Unauthorized Flow"
		}`)

		req, _ := http.NewRequest("POST", "/flows", requestBody)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		unauthRouter.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
}

func TestGetUserFlows_Integration(t *testing.T) {
	router, db, userID := setupTestEnvironment()
	defer db.Migrator().DropTable(&models.Flow{}, &models.User{})

	flows := []models.Flow{
		{ID: uuid.New(), Title: "Flow 1", UserID: userID},
		{ID: uuid.New(), Title: "Flow 2", UserID: userID},
	}
	for _, flow := range flows {
		db.Create(&flow)
	}

	t.Run("Success", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/flows", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response []models.Flow
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Len(t, response, 2)
	})

	t.Run("Unauthorized", func(t *testing.T) {
		unauthRouter := gin.Default()
		unauthRouter.Use(func(c *gin.Context) {
			c.Set("db", db)
		})
		unauthRouter.GET("/flows", handlers.GetUserFlows)

		req, _ := http.NewRequest("GET", "/flows", nil)
		w := httptest.NewRecorder()
		unauthRouter.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("Empty List", func(t *testing.T) {
		newUser := models.User{
			ID:       uuid.New(),
			Email:    "new@example.com",
			Password: "password",
		}
		db.Create(&newUser)

		newUserRouter := gin.Default()
		newUserRouter.Use(func(c *gin.Context) {
			c.Set("db", db)
			c.Set("userID", newUser.ID.String())
		})
		newUserRouter.GET("/flows", handlers.GetUserFlows)

		req, _ := http.NewRequest("GET", "/flows", nil)
		w := httptest.NewRecorder()
		newUserRouter.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response []models.Flow
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Empty(t, response)
	})
}
func TestGetFlow_Integration(t *testing.T) {
	router, db := setupTestEnvironment()
	defer db.Migrator().DropTable(&models.Flow{}, &models.User{})

	flow := models.Flow{
		ID:     uuid.New(),
		Title:  "Test Get Flow",
		Level:  "advanced",
		UserID: getTestUserID(db),
	}
	db.Create(&flow)

	t.Run("Get Existing Flow", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/flows/"+flow.ID.String(), nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response models.Flow
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, flow.ID, response.ID)
		assert.Equal(t, "Test Get Flow", response.Title)
	})

	t.Run("Get Non-Existent Flow", func(t *testing.T) {
		nonExistentID := uuid.New()
		req, _ := http.NewRequest("GET", "/flows/"+nonExistentID.String(), nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

func TestUpdateFlow_Integration(t *testing.T) {
	router, db, userID := setupTestEnvironment()
	defer db.Migrator().DropTable(&models.Flow{}, &models.User{})

	flow := models.Flow{
		ID:     uuid.New(),
		Title:  "Original Title",
		Level:  "beginner",
		UserID: userID,
	}
	db.Create(&flow)

	t.Run("Success", func(t *testing.T) {
		requestBody := bytes.NewBufferString(`{
			"title": "Updated Title",
			"level": "advanced",
			"cover": "new-cover.jpg"
		}`)

		req, _ := http.NewRequest("PUT", "/flows/"+flow.ID.String(), requestBody)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response models.Flow
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Updated Title", response.Title)
		assert.Equal(t, "advanced", response.Level)

		var dbFlow models.Flow
		db.First(&dbFlow, flow.ID)
		assert.Equal(t, "Updated Title", dbFlow.Title)
	})

	t.Run("Not Found", func(t *testing.T) {
		nonExistentID := uuid.New()
		requestBody := bytes.NewBufferString(`{"title": "Updated"}`)

		req, _ := http.NewRequest("PUT", "/flows/"+nonExistentID.String(), requestBody)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("Unauthorized", func(t *testing.T) {
		otherUserFlow := models.Flow{
			ID:     uuid.New(),
			Title:  "Other User Flow",
			UserID: uuid.New(),
		}
		db.Create(&otherUserFlow)

		requestBody := bytes.NewBufferString(`{"title": "Try to update"}`)

		req, _ := http.NewRequest("PUT", "/flows/"+otherUserFlow.ID.String(), requestBody)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("Invalid Input", func(t *testing.T) {
		requestBody := bytes.NewBufferString(`{"invalid": "json"`)

		req, _ := http.NewRequest("PUT", "/flows/"+flow.ID.String(), requestBody)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestDeleteFlow_Integration(t *testing.T) {
	router, db, userID := setupTestEnvironment()
	defer db.Migrator().DropTable(&models.Flow{}, &models.User{})

	t.Run("Success", func(t *testing.T) {
		flow := models.Flow{
			ID:     uuid.New(),
			Title:  "Flow to Delete",
			UserID: userID,
		}
		db.Create(&flow)

		req, _ := http.NewRequest("DELETE", "/flows/"+flow.ID.String(), nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var dbFlow models.Flow
		result := db.First(&dbFlow, flow.ID)
		assert.Error(t, result.Error)
		assert.Equal(t, gorm.ErrRecordNotFound, result.Error)
	})

	t.Run("Not Found", func(t *testing.T) {
		nonExistentID := uuid.New()
		req, _ := http.NewRequest("DELETE", "/flows/"+nonExistentID.String(), nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("Unauthorized", func(t *testing.T) {
		otherUserFlow := models.Flow{
			ID:     uuid.New(),
			Title:  "Other User Flow",
			UserID: uuid.New(),
		}
		db.Create(&otherUserFlow)

		req, _ := http.NewRequest("DELETE", "/flows/"+otherUserFlow.ID.String(), nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

