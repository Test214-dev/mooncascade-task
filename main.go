package main

import (
	"fmt"
	"github.com/Test214-dev/mooncascade-task/controllers"
	_ "github.com/Test214-dev/mooncascade-task/docs"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"os"
)

// @title Swagger API description
// @version 1.0
// @description Sport event server

// @host localhost
// @BasePath http://localhost:8080/api
func main() {
	dbHost := os.Getenv("POSTGRES_HOST")
	dbName := os.Getenv("POSTGRES_DB")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPswd := os.Getenv("POSTGRES_PASSWORD")
	connString := fmt.Sprintf("host=%s port=5432 user=%s sslmode=disable dbname=%s password=%s", dbHost, dbUser, dbName, dbPswd)
	db, err := gorm.Open("postgres", connString)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			fmt.Printf("Unable to close db connection: %s", err)
		}
	}()

	r := gin.Default()

	ac := controllers.AthleteController{DB: db}
	tc := controllers.TimingController{DB: db}

	// set up routes
	timings := r.Group("/api/timings")
	athletes := r.Group("/api/athletes")

	athletes.POST("/", ac.HandleAthletePost)
	athletes.GET("/", ac.HandleAthleteList)
	athletes.GET("/:id", ac.HandleAthleteGet)

	timings.GET("/", tc.HandleTimingList)
	timings.POST("/", tc.HandleTimingPost)
	timings.GET("/:id", tc.HandleTimingGet)

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	if err := r.Run(); err != nil {
		panic(err)
	}
}
