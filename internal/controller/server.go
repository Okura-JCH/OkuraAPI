package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartServer(router *gin.Engine, si ServerInterface) {
	RegisterHandlers(router, si)

	router.POST("/users", func(c *gin.Context) {
		// ここにユーザーを追加するロジックを実装
		c.JSON(http.StatusOK, gin.H{"message": "User added successfully"})
	})

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "UP"})
	})

	router.Run(":8080")
}
