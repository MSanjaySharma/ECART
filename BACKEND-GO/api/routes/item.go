package routes

import (
	"ecommerce/api/controller"
	"ecommerce/infrastructure"
)

//ItemRoute -> Route for question module
type ItemRoute struct {
	Controller controller.ItemController
	Handler    infrastructure.GinRouter
}

//NewItemRoute -> initializes new choice rouets
func NewItemRoute(
	controller controller.ItemController,
	handler infrastructure.GinRouter,

) ItemRoute {
	return ItemRoute{
		Controller: controller,
		Handler:    handler,
	}
}

//Setup -> setups new choice Routes
func (i ItemRoute) Setup() {
	item := i.Handler.Gin.Group("/item") //Router group
	{
		item.GET("/list", i.Controller.GetItems)
		item.POST("/create", i.Controller.AddItem)
	}
}
