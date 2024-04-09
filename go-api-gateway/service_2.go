package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func runMicroserviceTwo() {
	router := gin.Default()
	router.GET("/api/v1/info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Microservice 2",
		})
	})

	log.Fatal(router.Run(":8082"))
}
