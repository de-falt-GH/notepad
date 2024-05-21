package user

import (
	"kursarbeit/api/my_jwt"
	user_storage "kursarbeit/storage/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

	s.log.Info(id)

	res := DetailUserResponse{
		Login: user.Login,
		Email: user.Email,
		Name:  user.Name,
		Info:  user.Info,
	}

	ctx.IndentedJSON(http.StatusOK, res)
}

func (s service) UpdateUser(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")
	id, err := my_jwt.ExtractID(tokenString)
	if err != nil {
		s.log.Error(err)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "parsing jwt id failed"})
		return
	}

	var req UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		s.log.Error(err)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 15)
	if err != nil {
		s.log.Error(err)
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "password hashing failed"})
		return
	}

	if err = s.storage.UpdateUser(ctx, &user_storage.UpdateUserRequest{
		Id:           id,
		Login:        req.Login,
		PasswordHash: string(passwordHash),
		Email:        req.Email,
		Name:         req.Name,
		Info:         req.Info,
	}); err != nil {
		s.log.Error(err)
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "updating user failed"})
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

	res, err := s.storage.AddNote(ctx, &user_storage.AddNoteRequest{
		UserId: id,
		Name:   req.Name,
		Data:   req.Data,
		Public: req.Public,
	})
	if err != nil {
		s.log.Error(err)
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "adding note to db failed"})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"noteId": res.NoteId})
}

func (s service) UpdateNote(ctx *gin.Context) {
	var req UpdateNoteRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		s.log.Error(err)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
		return
	}

	if err := s.storage.UpdateNote(ctx, &user_storage.UpdateNoteRequest{
		Id:     req.Id,
		Name:   req.Name,
		Data:   req.Data,
		Public: req.Public,
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
		ctx.IndentedJSON(400, gin.H{"error": "invalid note id"})
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
	noteId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		s.log.Error(err)
		ctx.IndentedJSON(400, gin.H{"error": "invalid note id"})
		return
	}

	err = s.storage.DeleteNote(ctx, &user_storage.DeleteNoteRequest{Id: noteId})
	if err != nil {
		s.log.Error(err)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "database request failed"})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"msg": "note deleted successfully"})
}

func (s service) ListPublicNotes(ctx *gin.Context) {
	var req ListNotesRequest
	err := ctx.BindQuery(&req)
	if err != nil {
		s.log.Error(err)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
		return
	}

	if res, err := s.storage.ListPublicNotes(ctx, &user_storage.ListNotesRequest{
		Skip:  req.Skip,
		Limit: req.Limit,
	}); err != nil {
		s.log.Error(err)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
		return
	} else {
		ctx.IndentedJSON(http.StatusOK, res)
	}
}

func (s service) ListPrivateNotes(ctx *gin.Context) {
	var req ListNotesRequest
	err := ctx.BindQuery(&req)
	if err != nil {
		s.log.Error(err)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
		return
	}

	tokenString := ctx.GetHeader("Authorization")
	userId, err := my_jwt.ExtractID(tokenString)
	if err != nil {
		s.log.Error(err)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "parsing jwt id failed"})
		return
	}

	if res, err := s.storage.ListPrivateNotes(ctx, &user_storage.ListNotesRequest{
		UserId: userId,
		Skip:   req.Skip,
		Limit:  req.Limit,
	}); err != nil {
		s.log.Error(err)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
		return
	} else {
		ctx.IndentedJSON(http.StatusOK, res)
	}
}
