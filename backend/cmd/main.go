package main

import (
	"fmt"
	"net/http"

	"github.com/ProgramadoresSemPatria/Team-9/internal/config"
	"github.com/ProgramadoresSemPatria/Team-9/internal/database/connection"
	"github.com/ProgramadoresSemPatria/Team-9/internal/database/migrations"
	"github.com/ProgramadoresSemPatria/Team-9/internal/handlers"
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

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.POST("/register", handlers.CreateUserHandler)
	r.POST("/login", handlers.LoginHandler)
	authorized := r.Group("/", handlers.AuthMiddleware())
	{
		authorized.GET("/profile", handlers.ProfileHandler)
	}

	http.ListenAndServe(fmt.Sprintf(":%s", config.GetServerPort()), r)

}
