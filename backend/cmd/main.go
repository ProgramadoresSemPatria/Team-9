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
		AllowOrigins: []string{corsOrigin},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Accept",
			"Authorization",
			"X-Requested-With",
			"Access-Control-Allow-Headers",
		},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// Public routes (no authentication required)
	r.POST("/register", handlers.CreateUserHandler)
	r.POST("/login", handlers.LoginHandler)

	// Authenticated routes group (all routes below require valid JWT)
	authGroup := r.Group("/")
	authGroup.Use(handlers.AuthMiddleware()) // Apply authentication middleware to all routes in this group
	{
		// User profile routes
		authGroup.GET("/profile", handlers.ProfileHandler)

		// Flow management routes
		authGroup.POST("/flows", handlers.CreateFlow)
		authGroup.GET("/flows", handlers.GetUserFlows)

		// Flow-specific routes (operate on a single flow)
		flowRoutes := authGroup.Group("/flows/:id") // :id = flow ID parameter
		{
			flowRoutes.GET("", handlers.GetFlow)
			flowRoutes.PUT("", handlers.UpdateFlow)
			flowRoutes.DELETE("", handlers.DeleteFlow)

			// Workout day routes under a flow
			flowRoutes.POST("/workout-days", handlers.CreateWorkoutDay)
			flowRoutes.GET("/workout-days", handlers.GetWorkoutDaysByFlow)
		}

		// Workout day-specific routes (operate on a single workout day)
		workoutDayRoutes := authGroup.Group("/workout-days/:id") // :id = workout day ID
		{
			workoutDayRoutes.GET("", handlers.GetWorkoutDay)
			workoutDayRoutes.PUT("", handlers.UpdateWorkoutDay)
			workoutDayRoutes.DELETE("", handlers.DeleteWorkoutDay)

			// Exercise routes under a workout day
			workoutDayRoutes.POST("/exercises", handlers.CreateExercise)
			workoutDayRoutes.GET("/exercises", handlers.GetExercisesByWorkoutDay)
		}

		// Exercise-specific routes (operate on a single exercise)
		exerciseRoutes := authGroup.Group("/exercises/:id") // :id = exercise ID
		{
			exerciseRoutes.GET("", handlers.GetExercise)
			exerciseRoutes.PUT("", handlers.UpdateExercise)
			exerciseRoutes.DELETE("", handlers.DeleteExercise)
		}
	}

	http.ListenAndServe(fmt.Sprintf(":%s", config.GetServerPort()), r)

}
