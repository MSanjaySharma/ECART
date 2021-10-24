package routes

import (
	"ecommerce/api/controller"
	"ecommerce/infrastructure"
)

//OrderRoute -> Route for question module
type OrderRoute struct {
	Controller controller.OrderController
	Handler    infrastructure.GinRouter
}

//NewOrderRoute -> initializes new choice rouets
func NewOrderRoute(
	controller controller.OrderController,
	handler infrastructure.GinRouter,

) OrderRoute {
	return OrderRoute{
		Controller: controller,
		Handler:    handler,
	}
}

//Setup -> setups new choice Routes
func (o OrderRoute) Setup() {
	order := o.Handler.Gin.Group("/order") //Router group
	{
		order.GET("/list", o.Controller.GetOrders)
	}
}
