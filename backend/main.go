package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "Server is healthy.")
}

func main() {
	fmt.Println("Hello project Corporate Mundo")

	router := gin.Default()
	router.GET("/health", healthHandler)

	router.Run(":8030")
}
