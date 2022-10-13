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

// InsertOrderItems godoc
// @Tags orders
// @Description Create New Movie Data With Ordered Items.
// @ID create-new-orders
// @Accept json
// @Produce json
// @Param RequestBody body dto.NewOrderItemsRequest true "request body json"
// @Success 201 {object} dto.OrderItemsResponse
// @Router /orders [post]
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

// GetAllOrderItems godoc
// @Tags        orders
// @Description Retrieving All Order With Related Items
// @ID          get-all-orders
// @Produce     json
// @Success     200 {array} entity.Order
// @Router      /orders [get]
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

// UpdateOrderById godoc
// @Tags orders
// @Description Update existing order by ID and its items if needed
// @ID update-orders
// @Accept json
// @Produce json
// @Param orderID path int true "Order ID"
// @Success 201 {object} dto.OrderItemsResponse
// @Router /orders/{userID} [put]
func (o *orderRestHandler) UpdateOrderById(c *gin.Context) {
	orderID, err := helper.GetParamID(c, "orderID")

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"err":     "Invalid order ID",
		})
		return
	}

	var orderItemsRequest dto.NewOrderItemsRequest

	if err := c.ShouldBindJSON(&orderItemsRequest); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid JSON request",
			"err":     err.Error(),
		})
		return
	}

	updatedOrder, err := o.service.UpdateOrderItems(orderID, &orderItemsRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal server error",
			"err":     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedOrder)
}

// DeleteOrderById godoc
// @Tags orders
// @Description Update existing order by ID and its items if needed
// @ID delete-orders
// @Param orderID path int true "Order ID"
// @Success 201 {object} dto.OrderDeletedResponse "Order Deleted"
// @Router /orders/{userID} [delete]
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
