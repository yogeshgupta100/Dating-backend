package routes

import (
	"model/controller"
	"model/repository"
	"model/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CORSMiddleware() gin.HandlerFunc {
	allowedOrigins := []string{
		"https://hi.pokkoo.in",
		"https://www.hi.pokkoo.in",
	}

	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")

		// Check if the origin is allowed
		allow := false
		for _, o := range allowedOrigins {
			if o == origin {
				allow = true
				break
			}
		}

		if allow {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization, X-Requested-With")
			c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Header("Vary", "Origin")
		}

		// Preflight request
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	router.Use(CORSMiddleware())

	// Health check endpoint
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"status":  "healthy",
		})
	})

	// Additional health check with database
	router.GET("/health", func(c *gin.Context) {
		// Test database connection
		sqlDB, err := db.DB()
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"status": "unhealthy",
				"error":  "database connection failed",
			})
			return
		}

		if err := sqlDB.Ping(); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"status": "unhealthy",
				"error":  "database ping failed",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":   "healthy",
			"database": "connected",
		})
	})

	// Initialize repositories
	stateRepo := repository.NewStateRepository(db)
	modelRepo := repository.NewModelRepository(db)
	faqRepo := repository.NewFAQRepository(db)
	globalPhoneRepo := repository.NewGlobalPhoneRepository(db)

	// Initialize services
	stateService := service.NewStateService(stateRepo, modelRepo)
	modelService := service.NewModelService(modelRepo)
	faqService := service.NewFAQService(faqRepo)
	globalPhoneService := service.NewGlobalPhoneService(globalPhoneRepo)

	// Initialize controllers
	stateController := controller.NewStateController(stateService)
	modelController := controller.NewModelController(modelService)
	faqController := controller.NewFAQController(faqService)
	globalPhoneController := controller.NewGlobalPhoneController(globalPhoneService)

	// State routes
	router.POST("/states", stateController.CreateState)
	router.GET("/states", stateController.GetAllStates)
	router.GET("/states/:id", stateController.GetStateByID)
	router.GET("/states/slug/:slug", stateController.GetStateBySlug)
	router.GET("/states/:id/models", stateController.GetModelsByStateID)
	router.PUT("/states/:id", stateController.UpdateState)
	router.DELETE("/states/:id", stateController.DeleteState)

	// Model routes
	router.POST("/models", modelController.CreateModel)
	router.GET("/models", modelController.GetAllModels)
	router.GET("/models/:id", modelController.GetModelByID)
	router.PUT("/models/:id", modelController.UpdateModel)
	router.DELETE("/models/:id", modelController.DeleteModel)
	router.GET("/models/slug/:slug", modelController.GetModelsBySlug)

	// FAQ routes
	router.GET("/faq", faqController.GetFAQ)
	router.POST("/faq", faqController.CreateOrUpdateFAQ)

	// Global Phone routes
	router.POST("/global-phone", globalPhoneController.CreateOrUpdateGlobalPhone)
	router.GET("/global-phone", globalPhoneController.GetGlobalPhone)

	return router
}
