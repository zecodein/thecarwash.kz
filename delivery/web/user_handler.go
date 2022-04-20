package web

import (
	"github.com/gin-gonic/gin"
	"github.com/zecodein/thecarwash.kz/domain"
)

type UserHandler struct {
	userUsecase domain.UserUsecase
}

const key string = "thecarwash-session"

func NewUserHandler(r *gin.Engine, us domain.UserUsecase) {
	handler := &UserHandler{
		userUsecase: us,
	}

	r.POST("/user/signup", handler.signUp)
}

func (u *UserHandler) signUp(c *gin.Context) {
	// TODO
}
