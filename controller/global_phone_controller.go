package controller

import (
	"net/http"

	"model/models"
	"model/service"

	"github.com/gin-gonic/gin"
)

type GlobalPhoneController struct {
	globalPhoneService *service.GlobalPhoneService
}

func NewGlobalPhoneController(globalPhoneService *service.GlobalPhoneService) *GlobalPhoneController {
	return &GlobalPhoneController{
		globalPhoneService: globalPhoneService,
	}
}

func (c *GlobalPhoneController) CreateOrUpdateGlobalPhone(ctx *gin.Context) {
	var globalPhone models.GlobalPhone
	if err := ctx.ShouldBindJSON(&globalPhone); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.globalPhoneService.CreateOrUpdateGlobalPhone(&globalPhone); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, globalPhone)
}

func (c *GlobalPhoneController) GetGlobalPhone(ctx *gin.Context) {
	globalPhone, err := c.globalPhoneService.GetGlobalPhone()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Global phone number not found"})
		return
	}

	ctx.JSON(http.StatusOK, globalPhone)
}
