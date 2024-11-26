package handler

import (
	"golang-chapter-37/implem-gin-shipping/helper"
	"golang-chapter-37/implem-gin-shipping/model"
	"golang-chapter-37/implem-gin-shipping/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ShippingHadler struct {
	Service service.AllService
	Log     *zap.Logger
}

func NewShippingHandler(service service.AllService, log *zap.Logger) ShippingHadler {
	return ShippingHadler{
		Service: service,
		Log:     log,
	}
}

func (shippingHadler *ShippingHadler) Create(ca *gin.Context) {

}

func (shippingHadler *ShippingHadler) GetAllShipping(c *gin.Context) {

}

func (shippingHadler *ShippingHadler) ShippingCost(c *gin.Context) {
	var data model.RequestDestination
	if err := c.ShouldBindQuery(&data); err != nil {
		helper.BadResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	cost, err := shippingHadler.Service.CustomerService.ShippingCost(data)
	if err != nil {
		helper.BadResponse(c, err.Error(), http.StatusBadRequest)
	}

	result := gin.H{
		"cost": cost,
	}

	helper.SuccessResponseWithData(c, "success", http.StatusOK, result)
}
