package session

import (
	"database/sql"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strconv"
	dto2 "test/dto"
	"test/dto/dto_in"
	"test/dto/dto_out"
	repository "test/repository"
	"time"
)

func (s sessionService) Login(
	context *gin.Context,
) {
	var err error
	var dtoIn dto_in.LoginRequest
	err = context.ShouldBindJSON(&dtoIn)
	if err != nil {
		context.JSON(http.StatusBadRequest, dto2.ResponseAPI{
			Status:  http.StatusBadRequest,
			Message: "Failed " + err.Error(),
		})
		return
	}

	pass := base64.StdEncoding.EncodeToString([]byte(dtoIn.Password))
	userModel := repository.UserModel{
		Username: sql.NullString{String: dtoIn.Username},
		Password: sql.NullString{String: pass},
	}

	userOnDb, err := s.userDao.GetUserForLogin(userModel)
	if err != nil {
		context.JSON(http.StatusInternalServerError, dto2.ResponseAPI{
			Status:  http.StatusInternalServerError,
			Message: "Failed - Internal Server Error",
		})
		return
	}

	if userOnDb.ID.Int64 == 0 {
		context.JSON(http.StatusBadRequest, dto2.ResponseAPI{
			Status:  http.StatusBadRequest,
			Message: "Failed - User Not Found ",
		})
		return
	}
	var jwtKey = []byte("secret_key")
	expiredAt := time.Now().Add(24 * time.Hour)

	tokenModel := repository.PayloadJWTToken{
		UserID: userOnDb.ID.Int64,
		Claims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiredAt),
			Subject:   strconv.Itoa(int(userOnDb.ID.Int64)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenModel.Claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		context.JSON(http.StatusInternalServerError, dto2.ResponseAPI{
			Status:  http.StatusInternalServerError,
			Message: "Failed - Cannot generate token",
		})
		return
	}

	result := dto_out.LoginResponse{
		Username: dtoIn.Username,
		Password: "***",
		Name:     userOnDb.Name.string,
	}

	context.JSON(http.StatusOK, dto2.ResponseAPI{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    result,
		Token:   tokenString,
	})
	return
}
