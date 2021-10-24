package main

import (
	"ecommerce/api/controller"
	"ecommerce/api/repository"
	"ecommerce/api/routes"
	"ecommerce/api/service"
	"ecommerce/infrastructure"
	"ecommerce/models"
)

func init() {
	infrastructure.LoadEnv()
}

func main() {

	router := infrastructure.NewGinRouter()
	db := infrastructure.NewDatabase()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	userRoute := routes.NewUserRoute(userController, router)
	userRoute.Setup()

	itemRepository := repository.NewItemRepository(db)
	itemService := service.NewItemService(itemRepository)
	itemController := controller.NewItemController(itemService)
	itemRoute := routes.NewItemRoute(itemController, router)
	itemRoute.Setup()

	cartRepository := repository.NewCartRepository(db)
	cartService := service.NewCartService(cartRepository)
	cartController := controller.NewCartController(cartService)
	cartRoute := routes.NewCartRoute(cartController, router)
	cartRoute.Setup()

	orderRepository := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepository)
	orderController := controller.NewOrderController(orderService)
	orderRoute := routes.NewOrderRoute(orderController, router)
	orderRoute.Setup()

	db.DB.AutoMigrate(&models.Item{}, &models.Cart{}, &models.Order{}, &models.User{})

	router.Gin.Run(":7331")
}
