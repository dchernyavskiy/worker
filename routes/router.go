package routes

import (
	"github.com/gin-gonic/gin"
	"worker/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("views/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	clientRoutes(r)
	paymentRoutes(r)
	requestRoutes(r)
	serviceRoutes(r)
	providerRoutes(r)

	return r
}

func clientRoutes(r *gin.Engine) {
	userController := controllers.UserController{}

	// Endpoint to render the client_all.html template
	r.GET("/clients", userController.ShowClients)
	r.GET("/clients/create", userController.ShowCreateClient)
	r.GET("/clients/:id", userController.ShowClientDetails)
	r.GET("/clients/:id/update", userController.ShowUpdateClient)

	api := r.Group("/api")
	{
		api.POST("/clients", userController.CreateClient)
		api.POST("/clients/delete/:id", userController.DeleteClient)
		api.POST("/clients/update", userController.UpdateClient)
	}
}

func paymentRoutes(r *gin.Engine) {
	paymentController := controllers.PaymentController{}

	api := r.Group("/api")
	{
		api.POST("/payment", paymentController.CreatePayment)
	}
}

func requestRoutes(r *gin.Engine) {
	requestController := controllers.RequestController{}

	api := r.Group("/api")
	{
		api.POST("/request", requestController.CreateRequest)
	}
}

func serviceRoutes(r *gin.Engine) {
	serviceController := controllers.ServiceController{}

	// Endpoint to render the client_all.html template
	r.GET("/services", serviceController.ShowServices)
	r.GET("/services/create", serviceController.ShowCreateService)
	r.GET("/services/:id", serviceController.ShowServiceDetails)
	r.GET("/services/:id/update", serviceController.ShowUpdateService)

	api := r.Group("/api")
	{
		api.POST("/services", serviceController.CreateService)
		api.POST("/services/delete/:id", serviceController.DeleteService)
		api.POST("/services/update", serviceController.UpdateService)
	}
}

func providerRoutes(r *gin.Engine) {
	providerController := controllers.ProviderController{}

	r.GET("/providers/:id", providerController.ShowProviderDetails)
}
