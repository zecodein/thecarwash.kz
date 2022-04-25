package web

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getWashingByID(c *gin.Context) {
	param := c.Param("id")
	washingID, err := strconv.ParseInt(param, 10, 64)

	if washingID <= 0 && err == nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	if err != nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	washing, err := h.WashingUsecase.GetByID(c, washingID)
	if err != nil {
		c.JSON(errorHandler(err), nil)
		return
	}

	c.JSON(http.StatusOK, washing)
}
