package controller

import (
	"log"
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

	// Use GetPostForm to handle empty strings properly
	if _, exists := ctx.GetPostForm("phone_number"); exists {
		model.PhoneNumber = ctx.PostForm("phone_number")
	}
	if _, exists := ctx.GetPostForm("description"); exists {
		model.Description = ctx.PostForm("description")
	}
	if _, exists := ctx.GetPostForm("name"); exists {
		model.Name = ctx.PostForm("name")
	}
	if _, exists := ctx.GetPostForm("heading"); exists {
		model.Heading = ctx.PostForm("heading")
	}
	if _, exists := ctx.GetPostForm("profile_img"); exists {
		model.ProfileImg = ctx.PostForm("profile_img")
	}
	if _, exists := ctx.GetPostForm("banner_img"); exists {
		model.BannerImg = ctx.PostForm("banner_img")
	}
	if _, exists := ctx.GetPostForm("seo_title"); exists {
		model.SEOTitle = ctx.PostForm("seo_title")
	}
	if _, exists := ctx.GetPostForm("seo_desc"); exists {
		model.SEODesc = ctx.PostForm("seo_desc")
	}

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

	// Check if phone_number is present in form data (even if empty)
	if _, exists := ctx.GetPostForm("phone_number"); exists {
		phoneNumber := ctx.PostForm("phone_number")
		log.Printf("Updating phone_number for model %d: '%s' (exists: %v)", id, phoneNumber, exists)
		model.PhoneNumber = phoneNumber
	} else {
		log.Printf("phone_number field not present in form data for model %d", id)
	}

	// Check if description is present in form data (even if empty)
	if _, exists := ctx.GetPostForm("description"); exists {
		model.Description = ctx.PostForm("description")
	}

	// Check if name is present in form data (even if empty)
	if _, exists := ctx.GetPostForm("name"); exists {
		model.Name = ctx.PostForm("name")
	}

	if v := ctx.PostForm("heading"); v != "" {
		model.Heading = v
		// Generate new slug from updated heading
		model.Slug = models.GenerateSlug(v)
	}

	// Check if profile_img is present in form data (even if empty)
	if _, exists := ctx.GetPostForm("profile_img"); exists {
		model.ProfileImg = ctx.PostForm("profile_img")
	}

	// Check if banner_img is present in form data (even if empty)
	if _, exists := ctx.GetPostForm("banner_img"); exists {
		model.BannerImg = ctx.PostForm("banner_img")
	}

	// Check if seo_title is present in form data (even if empty)
	if _, exists := ctx.GetPostForm("seo_title"); exists {
		model.SEOTitle = ctx.PostForm("seo_title")
	}

	// Check if seo_desc is present in form data (even if empty)
	if _, exists := ctx.GetPostForm("seo_desc"); exists {
		model.SEODesc = ctx.PostForm("seo_desc")
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

	log.Printf("GetModelsBySlug: Searching for slug: '%s'", slug)

	models, err := c.modelService.GetModelsBySlug(slug)
	if err != nil {
		log.Printf("GetModelsBySlug: Error getting models for slug '%s': %v", slug, err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("GetModelsBySlug: Found %d models for slug '%s'", len(models), slug)
	for i, model := range models {
		log.Printf("GetModelsBySlug: Model %d - ID: %d, PhoneNumber: '%s', StateID: %d",
			i, model.ID, model.PhoneNumber, model.StateID)
	}

	if len(models) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No models found with the specified slug"})
		return
	}

	ctx.JSON(http.StatusOK, models)
}

// Debug endpoint to check database state
func (c *ModelController) DebugModels(ctx *gin.Context) {
	log.Printf("DebugModels: Checking all models in database")

	models, err := c.modelService.GetAllModels()
	if err != nil {
		log.Printf("DebugModels: Error getting all models: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("DebugModels: Found %d total models", len(models))
	for i, model := range models {
		log.Printf("DebugModels: Model %d - ID: %d, Slug: '%s', PhoneNumber: '%s', StateID: %d",
			i, model.ID, model.Slug, model.PhoneNumber, model.StateID)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"total_models": len(models),
		"models":       models,
	})
}
