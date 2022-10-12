package rest

import (
	"assignment-simple-rest-api/dto"
	"assignment-simple-rest-api/helper"
	"assignment-simple-rest-api/service"
	"fmt"
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
			"message": err.Error(),
			"err":     "INTERNAL_SERVER_ERROR",
		})
		return
	}

	c.JSON(http.StatusCreated, newOrder)
}

func (o *orderRestHandler) InsertOrderItems(c *gin.Context) {
	var orderItemsRequest dto.NewOrderItemsRequest

	if err := c.ShouldBindJSON(&orderItemsRequest); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid JSON request",
			"err":     err.Error(),
		})
		return
	}

	newOrderItems, err := o.service.InsertOrderItems(&orderItemsRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal server error",
			"err":     err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, newOrderItems)
}

func (o *orderRestHandler) GetAllOrderItems(c *gin.Context) {
	orderItems, err := o.service.GetAllOrderItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal server error",
			"err":     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, orderItems)
}

func (o *orderRestHandler) DeleteOrderByID(c *gin.Context) {
	orderID, err := helper.GetParamID(c, "orderID")

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"err":     "Invalid order ID",
		})
		return
	}

	if _, err := o.service.DeleteOrderByID(orderID); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal server error",
			"err":     err.Error(),
		})
		return
	}

	deletedMessage := fmt.Sprintf("Order with ID %d has been deleted", orderID)
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": deletedMessage,
	})
}
