package user

import (
	storage "kursarbeit/storage/user"

	auth "kursarbeit/api/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type service struct {
	storage storage.Storage
	log     *zap.SugaredLogger
}

func (s *service) SetRoutes(rg *gin.RouterGroup) {
	user := rg.Group("").Use(auth.Auth())

	user.GET("/profile", s.DetailUser)
	user.POST("/profile", s.UpdateUser)

	user.GET("/notes/private", s.ListPrivateNotes)
	user.POST("/notes", s.AddNote)
	user.GET("/notes/:id", s.DetailNote)
	user.PUT("/notes/:id", s.UpdateNote)
	user.DELETE("/notes/:id", s.DeleteNote)
}

func NewService(storage storage.Storage, log *zap.SugaredLogger) service {
	return service{storage: storage, log: log}
}
