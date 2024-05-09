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

	user.GET("/note/:id", s.DetailNote)
	user.POST("/note", s.AddNote)
	user.PUT("/note", s.UpdateNote)
	user.DELETE("/note", s.DeleteNote)
	user.POST("/note/list", s.ListNotes)

}

func NewService(storage storage.Storage, log *zap.SugaredLogger) service {
	return service{storage: storage, log: log}
}
