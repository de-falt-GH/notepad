package api

import (
	commonStorage "kursarbeit/storage/common"
	userStorage "kursarbeit/storage/user"
	"os"

	commonService "kursarbeit/api/service/common"
	userService "kursarbeit/api/service/user"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Serve(commonStorage commonStorage.Storage, userStorage userStorage.Storage, log *zap.SugaredLogger) {
	commonService := commonService.NewService(commonStorage, log)
	userService := userService.NewService(userStorage, log)

	router := gin.Default()
	commonService.SetRoutes(router.Group(""))
	userService.SetRoutes(router.Group(""))

	port := os.Getenv("BACKEND_PORT")
	router.Run("127.0.0.1:" + port)
}
