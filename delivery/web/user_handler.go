package web

import (
	"net/http"

	"github.com/gin-contrib/sessions"
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

	user.UserID, err = u.UserUsecase.Create(c, user)
	if err != nil {
		c.JSON(errorHandler(err), nil)
		return
	}

	user.ConfirmPassword = ""
	c.JSON(http.StatusCreated, user)
}

func (u *Handler) signIn(c *gin.Context) {
	user := &domain.User{}
	session := sessions.Default(c)

	if getSession(session) != 0 {
		c.JSON(http.StatusConflict, user)
		return
	}

	err := c.BindJSON(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, user)
		return
	}

	usr, err := u.UserUsecase.GetByNumber(c, user.Number)
	if err != nil {
		c.JSON(errorHandler(err), nil)
		return
	}

	if !u.UserUsecase.CheckPassword(c, usr.Password, user.Password) {
		c.JSON(http.StatusForbidden, user)
		return
	}

	err = setSession(session, usr.UserID)
	if err != nil {
		c.JSON(errorHandler(err), nil)
		return
	}

	c.JSON(http.StatusOK, usr)
}

func (u *Handler) signOut(c *gin.Context) {
	session := sessions.Default(c)

	if getSession(session) == 0 {
		c.JSON(http.StatusForbidden, nil)
		return
	}

	err := deleteSession(session)
	if err != nil {
		c.JSON(errorHandler(err), nil)
		return
	}

	c.JSON(http.StatusOK, nil)
}
