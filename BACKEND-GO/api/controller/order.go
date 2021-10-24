package controller

import (
	"ecommerce/api/service"
	"ecommerce/models"
	"ecommerce/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

//OrderController -> OrderController
type OrderController struct {
	service service.OrderService
}

//NewOrderController : NewOrderController
func NewOrderController(s service.OrderService) OrderController {
	return OrderController{
		service: s,
	}
}

// GetOrders : GetOrders controller
func (i OrderController) GetOrders(ctx *gin.Context) {
	var orders models.Order

	keyword := ctx.Query("keyword")

	data, total, err := i.service.FindAll(orders, keyword)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to find orders")
		return
	}
	respArr := make([]map[string]interface{}, 0)

	for _, n := range *data {
		resp := n.ResponseMap()
		respArr = append(respArr, resp)
	}

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Order result set",
		Data: map[string]interface{}{
			"rows":       respArr,
			"total_rows": total,
		}})
}
