package rest

import (
	"assignment-simple-rest-api/database"
	"assignment-simple-rest-api/docs"
	"assignment-simple-rest-api/repository/order_repository/order_pg"
	"assignment-simple-rest-api/service"
	"fmt"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

const port = "127.0.0.1:8080"

func StartApp() {
	database.InitializeDB()
	db := database.GetDB()

	// ! Inject Dependencies
	orderRepo := order_pg.NewOrderPg(db)
	orderService := service.NewOrderService(orderRepo)
	orderHandler := NewOrderRestHandler(orderService)

	// ! Routing API with gin
	route := gin.Default()

	// ! Swagger Routing
	docs.SwaggerInfo.Title = "Simple Order-Items API"
	docs.SwaggerInfo.Description = "Ini adalah API dengan pattern DDD hexagonal architecture, untuk membuat order dengan many items."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "127.0.0.1:8080"
	// docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Schemes = []string{"http"}

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	orderRoute := route.Group("/orders")
	{
		orderRoute.GET("/", orderHandler.GetAllOrderItems)
		orderRoute.POST("/", orderHandler.InsertOrderItems)
		orderRoute.PUT("/:orderID", orderHandler.UpdateOrderById)
		orderRoute.DELETE("/:orderID", orderHandler.DeleteOrderByID)
	}

	fmt.Println("Server is running on port", port)
	route.Run(port)
}
