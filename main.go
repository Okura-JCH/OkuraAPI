package main

import (
	"github.com/Okura-JCH/OkuraAPI.git/internal/controller"
	"github.com/Okura-JCH/OkuraAPI.git/pkg/db"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// データベース接続
	db := db.ConnectDB()
	defer db.Close()

	handler := controller.NewEndpointHandler(db)

	controller.StartServer(router, handler)
}
