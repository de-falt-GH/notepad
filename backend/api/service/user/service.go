package user

import (
	storage "kursarbeit/storage/user"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type service struct {
	storage storage.Storage
	log     *zap.SugaredLogger
}

func (s *service) SetRoutes(rg *gin.RouterGroup) {
	rg.POST("/login", s.postAuthorize)
	rg.POST("/register", s.postRegister)
}

func NewService(storage storage.Storage, log *zap.SugaredLogger) service {
	return service{storage: storage, log: log}
}
