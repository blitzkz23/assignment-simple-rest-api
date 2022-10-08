package rest

import (
	"assignment-simple-rest-api/dto"
	"assignment-simple-rest-api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type orderRestHandler struct {
	service service.OrderService
}

func NewOrderRestHandler(service service.OrderService) *orderRestHandler {
	return &orderRestHandler{
		service: service,
	}
}

func (o *orderRestHandler) InsertOrder(c *gin.Context) {
	orders, err := o.service.GetAllOrder()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal server error",
			"err":     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (o *orderRestHandler) GetAllOrder(c *gin.Context) {
	var orderRequest dto.NewOrderRequest

	if err := c.ShouldBindJSON(&orderRequest); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid JSON request",
			"err":     err.Error(),
		})
		return
	}

	newOrder, err := o.service.InsertOrder(&orderRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal server error",
			"err":     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, newOrder)
}
