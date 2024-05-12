package common

import (
	c_storage "kursarbeit/storage/common"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type service struct {
	storage c_storage.Storage
	log     *zap.SugaredLogger
}

func (s *service) SetRoutes(rg *gin.RouterGroup) {
	common := rg.Group("")

	common.POST("/login", s.postAuthorize)
	common.POST("/register", s.postRegister)
}

func NewService(storage c_storage.Storage, log *zap.SugaredLogger) service {
	return service{storage: storage, log: log}
}
