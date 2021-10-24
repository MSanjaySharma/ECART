package routes

import (
	"ecommerce/api/controller"
	"ecommerce/api/middlewares"
	"ecommerce/infrastructure"
)

//CartRoute -> Route for question module
type CartRoute struct {
	Controller controller.CartController
	Handler    infrastructure.GinRouter
}

//NewCartRoute -> initializes new choice rouets
func NewCartRoute(
	controller controller.CartController,
	handler infrastructure.GinRouter,

) CartRoute {
	return CartRoute{
		Controller: controller,
		Handler:    handler,
	}
}

//Setup -> setups new choice Routes
func (c CartRoute) Setup() {
	cart := c.Handler.Gin.Group("/cart") //Router group

	cart.Use(middlewares.EnforceAuthenticatedMiddleware())
	{
		cart.GET("/add", c.Controller.AddItemCart)
		//cart.POST("/:id/complete", c.Controller.MakeOrder)
	}
	{
		cart.GET("/list", c.Controller.GetCarts)
	}
}
