package web

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/zecodein/thecarwash.kz/domain"
)

type Handler struct {
	UserUsecase    domain.UserUsecase
	WashingUsecase domain.WashingUsecase
}

const Key string = "thecarwash-session"

func NewHandler(r *gin.Engine, h *Handler) {
	// * User handlers
	r.POST("/user/signup", h.signUp)
	r.POST("/user/signin", h.signIn)
	r.GET("/user/signout", h.signOut)
	// TODO update username, user number, user password

	// * Washing handlers
	r.POST("/washing/create", h.createWashing)
	r.GET("/washing/:id", h.getWashingByID)
}

func (h *Handler) createWashing(c *gin.Context) {
	userID := getSession(sessions.Default(c))
	if userID <= 0 {
		c.JSON(http.StatusForbidden, nil)
		return
	}

	// TODO раскоментить проверку на доступ
	// access, err := h.UserUsecase.GetAccess(c, userID)
	// if err != nil {
	// 	c.JSON(errorHandler(err), nil)
	// 	return
	// }

	// if access != "admin" {
	// 	c.JSON(http.StatusForbidden, nil)
	// 	return
	// }

	washing := &domain.Washing{}
	err := c.BindJSON(washing)
	if err != nil {
		c.JSON(errorHandler(domain.ErrInvalidData), nil)
		return
	}

	washing.WashingID, err = h.WashingUsecase.Create(c, washing)
	if err != nil {
		c.JSON(errorHandler(err), nil)
		return
	}

	c.JSON(http.StatusCreated, washing)
}
