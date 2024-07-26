package main

import (
	"fmt"
	"jobPostGraphQl/driver"
	"jobPostGraphQl/handler"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

const defaultPort = "8081"

func main() {
	engine := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	if err := driver.DatabaseConnection(); err != nil {
		fmt.Println("Failed to connect to DB", err)
		return
	}

	if err := driver.MigrateDB(); err != nil {
		fmt.Println("Failed to create to tables", err)
		return
	}

	engine.POST("/query", handler.GraphqlHandler())
	engine.GET("/", handler.PlaygroundHandler())

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	engine.Run(":" + port)
}
