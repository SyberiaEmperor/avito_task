package handler

import (
	"net/http"
	"strconv"

	"github.com/SyberiaEmperor/avito_task/models"
	"github.com/gin-gonic/gin"
)

const (
	accountId = "id"
)

func (h *Handler) getAccountInfo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param(accountId))

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Incorrect id")
		return
	}

	res, err := h.service.GetAccountInfo(id)

	if err != nil {
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"balance": res,
	})
}

func (h *Handler) deposit(c *gin.Context) {
	var req models.AccountRequest

	if err := c.BindJSON(&req); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err := h.service.Deposit(req)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": req.ID,
	})
}

func (h *Handler) debit(c *gin.Context) {

}

func (h *Handler) transfer(c *gin.Context) {
	var req models.TransferRequest

	if err := c.BindJSON(&req); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err := h.service.Transfer(req)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"sender_id":req.SenderID,
		"receiver_id":req.ReceiverID,
	})
}
