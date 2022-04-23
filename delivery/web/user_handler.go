package web

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zecodein/thecarwash.kz/domain"
)

func (u *Handler) signUp(c *gin.Context) {
	user := &domain.User{}

	err := c.BindJSON(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, user)
		return
	}

	fmt.Println(user)

	user.UserID, err = u.UserUsecase.Create(c, user)
	if err != nil {
		c.JSON(errorHandler(err), user)
		return
	}

	user.ConfirmPassword = ""
	c.JSON(http.StatusCreated, user)
}

func (u *Handler) signIn(c *gin.Context) {
	user := &domain.User{}

	err := c.BindJSON(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, user)
		return
	}

	usr, err := u.UserUsecase.GetByNumber(c, user.Number)
	if err != nil {
		c.JSON(errorHandler(err), user)
		return
	}

	if !u.UserUsecase.CheckPassword(c, usr.Password, user.Password) {
		c.JSON(http.StatusForbidden, user)
		return
	}
	// TODO set session
	c.JSON(http.StatusOK, usr)
}
