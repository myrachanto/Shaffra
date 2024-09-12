package routes

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/myrachanto/user/src/api/users"
	"github.com/swaggo/swag/example/basic/docs"
	"go.mongodb.org/mongo-driver/mongo"
)

var wg sync.WaitGroup

func ApiLoader(mongodb *mongo.Database) {
	ginApiServer(mongodb)
}
func ginApiServer(mongodb *mongo.Database) {
	u := users.NewUserController(users.NewUserService(users.NewUserRepo(mongodb)))
	docs.SwaggerInfo.BasePath = "/"
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.Default())

	router.POST("/users", concurrencyHandler(u.Create))
	router.GET("/users", concurrencyHandler(u.GetAll))
	router.GET("/users/:id", concurrencyHandler(u.GetOne))
	router.PUT("/users/:id", concurrencyHandler(u.Update))
	router.DELETE("/users/:id", concurrencyHandler(u.Delete))
	router.GET("/health", HealthCheck)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file in routes")
	}

	PORT := os.Getenv("HTTP_PORT")
	router.Run(":" + PORT)
}

// Concurrency middleware to log request time for Gin and user controller methods
func concurrencyHandler(handler func(c *gin.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		// Add a new goroutine to the WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done() // Mark this goroutine as done

			// Handle the actual request
			handler(c)

			// Log the time taken for the request
			elapsedTime := time.Since(startTime)
			log.Printf("Request took %s", elapsedTime)
		}()

		// Wait for all goroutines to finish
		wg.Wait()
	}
}
