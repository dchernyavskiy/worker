package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"worker/database"
	"worker/models"
)

type UserController struct{}

func (u *UserController) GetUsers(c *gin.Context) {
	var users []models.User
	database.DB.Find(&users)

	// Render the user_all.html template with the user data
	c.HTML(200, "user_all.html", gin.H{"Users": users})
}

func (u *UserController) CreateUser(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBind(&newUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&newUser)
	c.Redirect(http.StatusFound, "./../users")
}

func (u *UserController) ShowCreateUserForm(c *gin.Context) {
	c.HTML(200, "user_create.html", nil)
}

const (
	View   string = "view"
	Update        = "update"
	Delete        = "delete"
)

type UserVm struct {
	models.User
	action string
}

func (u *UserController) ViewUpdateDelete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	result := database.DB.FirstOrInit(&user, id)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}

	action := c.Query("action")

	vm := UserVm{
		User:   user,
		action: action,
	}

	c.HTML(200, "user_one.html", vm)
}
