package controller

import "github.com/gin-gonic/gin"

func StartServer(si ServerInterface) {
	g := gin.Default()

	RegisterHandlers(g, si)

	g.Run()
}
