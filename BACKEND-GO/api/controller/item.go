package controller

import (
	"ecommerce/api/service"
	"ecommerce/models"
	"ecommerce/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

//ItemController -> ItemController
type ItemController struct {
	service service.ItemService
}

//NewItemController : NewItemController
func NewItemController(s service.ItemService) ItemController {
	return ItemController{
		service: s,
	}
}

// GetItems : GetItems controller
func (i ItemController) GetItems(ctx *gin.Context) {
	var items models.Item

	keyword := ctx.Query("keyword")

	data, total, err := i.service.FindAll(items, keyword)

	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to find items")
		return
	}
	respArr := make([]map[string]interface{}, 0)

	for _, n := range *data {
		resp := n.ResponseMap()
		respArr = append(respArr, resp)
	}

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Item result set",
		Data: map[string]interface{}{
			"rows":       respArr,
			"total_rows": total,
		}})
}

// AddItem : AddItem controller
func (i *ItemController) AddItem(ctx *gin.Context) {
	var item models.Item
	ctx.ShouldBindJSON(&item)

	if item.Name == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Name is required")
		return
	}

	err := i.service.Save(item)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Failed to create item")
		return
	}
	util.SuccessJSON(ctx, http.StatusCreated, "Successfully Created Item")
}
