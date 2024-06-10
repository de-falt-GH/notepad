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
	common.GET("/notes/public", s.ListPublicNotes)
	common.GET("/notes/:id", s.DetailNote)
}

func NewService(storage c_storage.Storage, log *zap.SugaredLogger) service {
	return service{storage: storage, log: log}
}
