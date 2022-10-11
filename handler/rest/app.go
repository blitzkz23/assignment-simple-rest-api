package rest

import (
	"assignment-simple-rest-api/database"
	"assignment-simple-rest-api/repository/order_repository/order_pg"
	"assignment-simple-rest-api/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

const port = ":8080"

func StartApp() {
	database.InitializeDB()
	db := database.GetDB()

	// ! Inject Dependencies
	orderRepo := order_pg.NewOrderPg(db)
	orderService := service.NewOrderService(orderRepo)
	orderHandler := NewOrderRestHandler(orderService)

	// ! Routing API with gin
	route := gin.Default()

	orderRoute := route.Group("/orders")
	{
		orderRoute.GET("/", orderHandler.GetAllOrderItems)
		orderRoute.POST("/", orderHandler.InsertOrder)
		orderRoute.DELETE("/:orderID", orderHandler.DeleteOrderByID)
	}

	fmt.Println("Server is running on port", port)
	route.Run(port)
}
