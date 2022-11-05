package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getAccountInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	// if !ok {
	// 	newErrorResponse(c, http.StatusInternalServerError, "id not found")
	// 	return
	// }

	res, err := h.service.GetAccountInfo(id)

	if err != nil {
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"balande": res,
	})
}

func (h *Handler) deposit(c *gin.Context) {

}

func (h *Handler) debit(c *gin.Context) {

}

func (h *Handler) transfer(c *gin.Context) {

}
