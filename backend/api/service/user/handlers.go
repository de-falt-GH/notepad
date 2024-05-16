package user

import (
	"kursarbeit/api/my_jwt"
	user_storage "kursarbeit/storage/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s service) DetailUser(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")
	id, err := my_jwt.ExtractID(tokenString)
	if err != nil {
		s.log.Error(err)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "parsing jwt id failed"})
		return
	}

	user, err := s.storage.DetailUser(ctx, &user_storage.DetailUserRequest{Id: id})
	if err != nil {
		s.log.Error(err)
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	res := DetailUserResponse{
		login: user.Login,
		email: user.Email,
		name:  user.Name,
		info:  user.Info,
	}

	ctx.IndentedJSON(http.StatusOK, res)
}

func (s service) UpdateUser(ctx *gin.Context) {
	var req UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		s.log.Error(err)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"msg": "user updated successfully"})
}

func (s service) AddNote(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")
	id, err := my_jwt.ExtractID(tokenString)
	if err != nil {
		s.log.Error(err)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "parsing jwt id failed"})
		return
	}

	var req AddNoteRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		s.log.Error(err)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
		return
	}

	if err := s.storage.AddNote(ctx, &user_storage.AddNoteRequest{
		UserId: id,
		Name:   req.name,
		Data:   req.data,
		Public: req.public,
	}); err != nil {
		s.log.Error(err)
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "adding note to db failed"})
		return
	}
}

func (s service) UpdateNote(ctx *gin.Context) {
	var req UpdateNoteRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		s.log.Error(err)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
		return
	}

	if err := s.storage.UpdateNote(ctx, &user_storage.UpdateNoteRequest{
		Id:     req.id,
		Name:   req.name,
		Data:   req.data,
		Public: req.public,
	}); err != nil {
		s.log.Error(err)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "updating note failed failed"})
		return
	}
}

func (s service) DetailNote(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		s.log.Error(err)
		ctx.IndentedJSON(400, gin.H{"error": "invalid order id"})
		return
	}

	if res, err := s.storage.DetailNote(ctx, &user_storage.DetailNoteRequest{
		Id: id,
	}); err != nil {
		s.log.Error(err)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "updating note failed failed"})
		return
	} else {
		ctx.IndentedJSON(http.StatusOK, res)
	}
}

func (s service) DeleteNote(ctx *gin.Context) {
	var req DeleteNoteRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		s.log.Error(err)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"msg": "note deleted successfully"})
}

func (s service) ListNotes(ctx *gin.Context) {
	var req ListNotesRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		s.log.Error(err)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
		return
	}

	tokenString := ctx.GetHeader("Authorization")
	id, err := my_jwt.ExtractID(tokenString)
	if err != nil {
		s.log.Error(err)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "parsing jwt id failed"})
		return
	}

	if res, err := s.storage.ListNotes(ctx, &user_storage.ListNotesRequest{
		UserId:   id,
		Skip:     req.skip,
		Limit:    req.limit,
		Personal: req.personal,
	}); err != nil {
		s.log.Error(err)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
		return
	} else {
		ctx.IndentedJSON(http.StatusOK, res)
	}
}
