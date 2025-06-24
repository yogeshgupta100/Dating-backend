package routes

import (
	"model/controller"
	"model/repository"
	"model/service"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		log.Printf("Request - Origin: %s, Method: %s, Path: %s, Headers: %v",
			origin, c.Request.Method, c.Request.URL.Path, c.Request.Header)

		allowedOrigins := map[string]bool{
			"https://dating-backend-wzzl.onrender.com": true,
			"https://pro.abellarora.com":               true,
		}

		setCORSHeaders := func() {
			if allowedOrigins[origin] {
				c.Header("Access-Control-Allow-Origin", origin)
			} else {
				log.Printf("Unrecognized origin: %s", origin)
				c.Header("Access-Control-Allow-Origin", "http://localhost:5173")
			}
			c.Header("Vary", "Origin")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, ngrok-skip-browser-warning, access-control-allow-headers, access-control-allow-methods, access-control-allow-origin")
			c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			log.Printf("CORS Headers: %v", c.Writer.Header())
		}

		setCORSHeaders()

		if c.Writer.Header().Get("Access-Control-Allow-Origin") == "" {
			log.Printf("CORS headers missing, reapplying - Path: %s, Status: %d", c.Request.URL.Path, c.Writer.Status())
			setCORSHeaders()
		}

		if c.Request.Method == http.MethodOptions {
			log.Printf("Handling OPTIONS preflight for %s", c.Request.URL.Path)
			c.Header("Cache-Control", "no-store, no-cache, must-revalidate")
			c.Header("Access-Control-Max-Age", "86400")
			setCORSHeaders()
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		defer func() {
			setCORSHeaders()
			if c.Writer.Status() >= 300 {
				log.Printf("Special response (redirect/error) - Path: %s, Method: %s, Status: %d, Headers: %v",
					c.Request.URL.Path, c.Request.Method, c.Writer.Status(), c.Writer.Header())
			}
			log.Printf("Response - Path: %s, Method: %s, Status: %d, Headers: %v",
				c.Request.URL.Path, c.Request.Method, c.Writer.Status(), c.Writer.Header())
		}()

		c.Next()
	}
}

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	router.Use(CORSMiddleware())

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

	// FAQ routes
	router.GET("/faq", faqController.GetFAQ)
	router.POST("/faq", faqController.CreateOrUpdateFAQ)

	// Global Phone routes
	router.POST("/global-phone", globalPhoneController.CreateOrUpdateGlobalPhone)
	router.GET("/global-phone", globalPhoneController.GetGlobalPhone)

	return router
}
