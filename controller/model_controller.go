package controller

import (
	"net/http"
	"strconv"

	"model/models"
	"model/service"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

type ModelController struct {
	modelService *service.ModelService
}

func NewModelController(modelService *service.ModelService) *ModelController {
	return &ModelController{
		modelService: modelService,
	}
}

func (c *ModelController) CreateModel(ctx *gin.Context) {
	var model models.Model

	// Parse form fields
	stateIDStr := ctx.PostForm("state_id")
	if stateIDStr != "" {
		id, _ := strconv.ParseUint(stateIDStr, 10, 32)
		model.StateID = uint(id)
	}
	model.PhoneNumber = ctx.PostForm("phone_number")
	model.Description = ctx.PostForm("description")
	model.Name = ctx.PostForm("name")
	model.Heading = ctx.PostForm("heading")
	model.ProfileImg = ctx.PostForm("profile_img")
	model.BannerImg = ctx.PostForm("banner_img")
	model.SEOTitle = ctx.PostForm("seo_title")
	model.SEODesc = ctx.PostForm("seo_desc")

	// Generate slug from heading
	if model.Heading != "" {
		model.Slug = models.GenerateSlug(model.Heading)
	}

	// Parse services as JSON array string
	services := ctx.PostForm("services")
	if services != "" {
		json.Unmarshal([]byte(services), &model.Services)
	}

	if err := c.modelService.CreateModel(&model); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, model)
}

func (c *ModelController) GetAllModels(ctx *gin.Context) {
	models, err := c.modelService.GetAllModels()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, models)
}

func (c *ModelController) GetModelByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	model, err := c.modelService.GetModelByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Model not found"})
		return
	}

	ctx.JSON(http.StatusOK, model)
}

func (c *ModelController) UpdateModel(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	model, err := c.modelService.GetModelByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Model not found"})
		return
	}

	// Parse form fields and update only if provided
	if v := ctx.PostForm("state_id"); v != "" {
		if sid, err := strconv.ParseUint(v, 10, 32); err == nil {
			model.StateID = uint(sid)
		}
	}
	if v := ctx.PostForm("phone_number"); v != "" {
		model.PhoneNumber = v
	}
	if v := ctx.PostForm("description"); v != "" {
		model.Description = v
	}
	if v := ctx.PostForm("name"); v != "" {
		model.Name = v
	}
	if v := ctx.PostForm("heading"); v != "" {
		model.Heading = v
		// Generate new slug from updated heading
		model.Slug = models.GenerateSlug(v)
	}
	if v := ctx.PostForm("profile_img"); v != "" {
		model.ProfileImg = v
	}
	if v := ctx.PostForm("banner_img"); v != "" {
		model.BannerImg = v
	}
	if v := ctx.PostForm("seo_title"); v != "" {
		model.SEOTitle = v
	}
	if v := ctx.PostForm("seo_desc"); v != "" {
		model.SEODesc = v
	}
	if v := ctx.PostForm("services"); v != "" {
		var services []string
		if err := json.Unmarshal([]byte(v), &services); err == nil {
			model.Services = services
		}
	}

	if err := c.modelService.UpdateModel(model); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, model)
}

func (c *ModelController) DeleteModel(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := c.modelService.DeleteModel(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Model deleted successfully"})
}

func (c *ModelController) GetModelsByHeading(ctx *gin.Context) {
	heading := ctx.Param("heading")
	if heading == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Heading parameter is required"})
		return
	}

	models, err := c.modelService.GetModelsByHeading(heading)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(models) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No models found with the specified heading"})
		return
	}

	ctx.JSON(http.StatusOK, models)
}

func (c *ModelController) GetModelsBySlug(ctx *gin.Context) {
	slug := ctx.Param("slug")
	if slug == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Slug parameter is required"})
		return
	}

	models, err := c.modelService.GetModelsBySlug(slug)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(models) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No models found with the specified slug"})
		return
	}

	ctx.JSON(http.StatusOK, models)
}
