package common

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"

	"kursarbeit/api/my_jwt"
	c_storage "kursarbeit/storage/common"
)

func (s *service) postRegister(ctx *gin.Context) {
	var req registerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid format"})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 15)
	if err != nil {
		s.log.Error(err)
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "password hashing failed"})
		return
	}

	id, err := s.storage.CreateUser(ctx, &c_storage.CreateUserRequest{
		Login:        req.Login,
		PasswordHash: string(passwordHash),
		Email:        req.Email,
		Name:         req.Name,
		Info:         req.Info,
	})
	if err != nil {
		s.log.Error(err)
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "saving to db failed"})
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = int64(time.Now().Add(720 * time.Hour).Unix()) // week
	claims["uid"] = id
	tokenString, err := token.SignedString(my_jwt.Salt)
	if err != nil {
		s.log.Error(err)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "unable to generate token string: " + err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"token": tokenString})
}

func (s *service) postAuthorize(ctx *gin.Context) {
	var req authorizeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		s.log.Error(err)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid format"})
		return
	}

	user, err := s.storage.DetailUser(ctx, &c_storage.DetailUserRequest{Login: req.Login})
	if err != nil || bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)) != nil {
		s.log.Error(err)
		ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = int64(time.Now().Add(720 * time.Hour).Unix()) // week
	claims["uid"] = user.Id
	tokenString, err := token.SignedString(my_jwt.Salt)
	if err != nil {
		s.log.Error(err)
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "unable to generate token string: " + err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"token": tokenString})
}

func (s service) ListPublicNotes(ctx *gin.Context) {
	var req ListNotesRequest
	err := ctx.BindQuery(&req)
	if err != nil {
		s.log.Error(err)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
		return
	}

	res, err := s.storage.ListPublicNotes(ctx, &c_storage.ListPublicNotesRequest{
		Search: req.Search,
		Skip:   req.Skip,
		Limit:  req.Limit,
	})
	if err != nil {
		s.log.Error(err)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
		return
	}

	notes := []NoteListItem{}
	for _, note := range res {
		notes = append(notes, NoteListItem{Id: note.Id, AuthorId: note.UserId, Name: note.Name, AuthorName: note.AuthorName, Updated: note.Updated})
	}

	ctx.IndentedJSON(http.StatusOK, ListPublicNotesResponse{Notes: notes})
}
