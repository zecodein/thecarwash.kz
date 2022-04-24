package web

import (
	"net/http"
	"strconv"

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

	r.GET("/washing/:id", h.washing)
}

func (h *Handler) washing(c *gin.Context) {
	param := c.Param("id")
	washingID, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		c.JSON(errorHandler(err), nil)
		return
	}

	c.JSON(http.StatusOK, washingID)
	// TODO
}
