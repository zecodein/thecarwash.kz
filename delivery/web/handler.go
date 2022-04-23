package web

import (
	"github.com/gin-gonic/gin"
	"github.com/zecodein/thecarwash.kz/domain"
)

type Handler struct {
	UserUsecase domain.UserUsecase
}

const Key string = "thecarwash-session"

func NewHandler(r *gin.Engine, h *Handler) {
	r.POST("/user/signup", h.signUp)
	r.POST("/user/signin", h.signIn)
	r.GET("/user/signout", h.signOut)
	// TODO update username, user number, user password
}
