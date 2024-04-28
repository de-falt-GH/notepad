package common

import "github.com/gin-gonic/gin"

func (s *service) SetRoutes(rg *gin.RouterGroup) {
	rg.POST("/login", s.postAuthorize)
	rg.POST("/register", s.postRegister)
}
