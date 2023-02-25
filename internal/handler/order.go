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
		//newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		log.Fatalf("Param error - invalid id: %s", err.Error())
	}

	order, err := h.service.GetOrderById(id)
	if err != nil {
		//newErrorResponse(c, http.StatusInternalServerError, err.Error())
		//return
	}

	c.JSON(http.StatusOK, order)

}

//func (h *Handler) createOrder(c *gin.Context) {
//
//	var input model.Order
//
//	//не читает массив items
//	if err := c.BindJSON(&input); err != nil {
//		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
//		log.Fatalf("Not valid data error: %s", err.Error())
//		/*newErrorResponse(c, http.StatusBadRequest, err.Error())
//		return*/
//	}
//
//	//fmt.Printf("data: %s\n", input)
//	id, err := h.service.CreateOrder(input)
//
//	if err != nil {
//		newErrorResponse(c, http.StatusInternalServerError, err.Error())
//		return
//	}
//	c.JSON(http.StatusOK, map[string]interface{}{
//		"id": id,
//	})
//}
