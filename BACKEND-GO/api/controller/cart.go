package controller

import (
	"ecommerce/api/service"
	"ecommerce/models"
	"ecommerce/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

//CartController -> CartController
type CartController struct {
	service service.CartService
}

//NewCartController : NewCartController
func NewCartController(s service.CartService) CartController {
	return CartController{
		service: s,
	}
}

// GetCarts : GetCarts controller
func (i CartController) GetCarts(ctx *gin.Context) {
	var carts models.Cart

	keyword := ctx.Query("keyword")

	data, total, err := i.service.FindAll(carts, keyword)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to find carts")
		return
	}
	respArr := make([]map[string]interface{}, 0)

	for _, n := range *data {
		resp := n.ResponseMap()
		respArr = append(respArr, resp)
	}

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Cart result set",
		Data: map[string]interface{}{
			"rows":       respArr,
			"total_rows": total,
		}})
}

// AddItemToCart : AddItemToCart controller
func (i *CartController) AddItemCart(ctx *gin.Context) {
	var cart models.Cart
	ctx.ShouldBindJSON(&cart)

	err := i.service.Save(cart)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to create item")
		return
	}
	util.SuccessJSON(ctx, http.StatusCreated, "Successfully Created Item")
}
