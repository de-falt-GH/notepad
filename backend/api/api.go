package api

import (
	commonStorage "kursarbeit/storage/common"
	userStorage "kursarbeit/storage/user"
	"net/http"
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

	router.Use(func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if ctx.Request.Method == "OPTIONS" {
			ctx.Writer.WriteHeader(http.StatusOK)
		}

		ctx.Next()
	})

	commonService.SetRoutes(router.Group(""))
	userService.SetRoutes(router.Group(""))

	router.Run(os.Getenv("BACKEND_HOST") + ":" + os.Getenv("BACKEND_PORT"))
}
