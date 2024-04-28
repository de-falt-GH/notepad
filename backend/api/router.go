package api

import (
	"kursarbeit/api/handlers/common"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()

	common.SetRoutes(router.Group(""))
	admin.SetRoutes(router.Group(""))
	manager.SetRoutes(router.Group(""))
	master.SetRoutes(router.Group(""))
	customer.SetRoutes(router.Group(""))

	router.Run("127.0.0.1:8080")
}
