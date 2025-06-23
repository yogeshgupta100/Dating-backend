package controller

import (
	"model/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FAQController struct {
	faqService *service.FAQService
}

func NewFAQController(faqService *service.FAQService) *FAQController {
	return &FAQController{faqService: faqService}
}

func (c *FAQController) GetFAQ(ctx *gin.Context) {
	faq, err := c.faqService.GetFAQ()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "FAQ not found"})
		return
	}
	ctx.JSON(http.StatusOK, faq)
}

func (c *FAQController) CreateOrUpdateFAQ(ctx *gin.Context) {
	var req struct {
		Name string `json:"name" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	faq, err := c.faqService.CreateOrUpdateFAQ(req.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, faq)
}
