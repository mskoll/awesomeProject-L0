package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// getOrderById обработка запроса на получение order
func (h *Handler) getOrderById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatalf("Param error - invalid id: %s", err.Error())
	}

	order, err := h.service.GetOrderById(id)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	c.JSON(http.StatusOK, order)

}
