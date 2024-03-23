package main

import (
	"net/http"

	"github.com/Okura-JCH/OkuraAPI.git/pkg/db"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// データベース接続
	db := db.ConnectDB()
	defer db.Close()

	// ルーティング設定
	router.POST("/users", func(c *gin.Context) {
		// ここにユーザーを追加するロジックを実装
		c.JSON(http.StatusOK, gin.H{"message": "User added successfully"})
	})

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "UP"})
	})

	router.Run(":8080")
}
