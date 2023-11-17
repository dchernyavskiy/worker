package routes

import (
	"github.com/gin-gonic/gin"
	"worker/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("views/*")

	userRoutes(r)

	return r
}

func userRoutes(r *gin.Engine) {
	userController := controllers.UserController{}

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	// Endpoint to render the user_all.html template
	r.GET("/users", userController.GetUsers)
	r.GET("/users/create", userController.ShowCreateUserForm)
	r.GET("/users/:id", userController.ViewUpdateDelete)

	api := r.Group("/api")
	{
		api.POST("/users", userController.CreateUser)
		// Add other API routes as needed
	}
}
