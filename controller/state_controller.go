package controller

import (
	"net/http"
	"strconv"

	"model/models"
	"model/service"

	"github.com/gin-gonic/gin"
)

type StateController struct {
	stateService *service.StateService
}

func NewStateController(stateService *service.StateService) *StateController {
	return &StateController{
		stateService: stateService,
	}
}

func (c *StateController) CreateState(ctx *gin.Context) {
	var state models.State
	if err := ctx.ShouldBindJSON(&state); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.stateService.CreateState(&state); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, state)
}

func (c *StateController) GetAllStates(ctx *gin.Context) {
	states, err := c.stateService.GetAllStates()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, states)
}

func (c *StateController) GetStateByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	state, err := c.stateService.GetStateByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "State not found"})
		return
	}

	ctx.JSON(http.StatusOK, state)
}

func (c *StateController) GetStateBySlug(ctx *gin.Context) {
	slug := ctx.Param("slug")
	if slug == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Slug is required"})
		return
	}

	state, err := c.stateService.GetStateBySlug(slug)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "State not found"})
		return
	}

	ctx.JSON(http.StatusOK, state)
}

func (c *StateController) GetModelsByStateID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	models, err := c.stateService.GetModelsByStateID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, models)
}

func (c *StateController) UpdateState(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var req struct {
		Name        string `json:"name"`
		Slug        string `json:"slug"`
		PhoneNumber string `json:"phone_number"`
		Heading     string `json:"heading"`
		SubHeading  string `json:"sub_heading"`
		Content     string `json:"content"`
		SEOTitle    string `json:"seo_title"`
		SEODesc     string `json:"seo_desc"`
		SEOKeyword  string `json:"seo_keyword"`
		FAQ         string `json:"faq"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	state, err := c.stateService.GetStateByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "State not found"})
		return
	}

	state.Name = req.Name
	state.Slug = req.Slug
	state.PhoneNumber = req.PhoneNumber
	state.Heading = req.Heading
	state.SubHeading = req.SubHeading
	state.Content = req.Content
	state.SEOTitle = req.SEOTitle
	state.SEODesc = req.SEODesc
	state.SEOKeyword = req.SEOKeyword
	state.FAQ = req.FAQ

	if err := c.stateService.UpdateState(state); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, state)
}

func (c *StateController) DeleteState(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := c.stateService.DeleteState(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "State deleted successfully"})
}
