package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/ProgramadoresSemPatria/Team-9/internal/config"
	"github.com/ProgramadoresSemPatria/Team-9/internal/database/connection"
	"github.com/ProgramadoresSemPatria/Team-9/internal/database/migrations"
	"github.com/ProgramadoresSemPatria/Team-9/internal/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	err = config.Load()
	if err != nil {
		panic(fmt.Sprintf("Failed to load configuration: %v", err))
	}

	db, err := connection.OpenConnection()
	if err != nil {
		panic(err)
	}

	migrations.RunMigrations(db)

	r := gin.Default()

	corsOrigin := os.Getenv("CORS")

	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{corsOrigin},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        12 * time.Hour,
	}))

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.POST("/register", handlers.CreateUserHandler)
	r.POST("/login", handlers.LoginHandler)

	authGroup := r.Group("/")
	authGroup.Use(handlers.AuthMiddleware())
	{
		authGroup.GET("/profile", handlers.ProfileHandler)

		authGroup.POST("/flows", handlers.CreateFlow)
		authGroup.GET("/flows", handlers.GetUserFlows)
		authGroup.GET("/flows/:id", handlers.GetFlow)
		authGroup.PUT("/flows/:id", handlers.UpdateFlow)
		authGroup.DELETE("/flows/:id", handlers.DeleteFlow)
	}

	http.ListenAndServe(fmt.Sprintf(":%s", config.GetServerPort()), r)

}
