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

func getTestUserID(db *gorm.DB) uuid.UUID {
	var user models.User
	db.Where("email = ?", "test@example.com").First(&user)
	return user.ID
}

